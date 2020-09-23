#include <memory>
#include <string>
#include <unordered_set>
#include <vector>

#include "src/carnot/planner/compiler/ast_visitor.h"
#include "src/carnot/planner/distributed/distributed_analyzer.h"
#include "src/carnot/planner/distributed/distributed_coordinator.h"
#include "src/carnot/planner/distributed/distributed_planner.h"
#include "src/carnot/planner/rules/rules.h"
namespace pl {
namespace carnot {
namespace planner {
namespace distributed {

StatusOr<std::unique_ptr<DistributedPlanner>> DistributedPlanner::Create() {
  std::unique_ptr<DistributedPlanner> planner(new DistributedPlanner());
  PL_RETURN_IF_ERROR(planner->Init());
  return planner;
}

Status DistributedPlanner::Init() { return Status::OK(); }
Status StitchPlan(DistributedPlan* distributed_plan) {
  DCHECK(distributed_plan);
  auto remote_carnot = distributed_plan->kelvin();
  DCHECK(remote_carnot);
  auto remote_node_id = remote_carnot->id();
  IR* remote_plan = remote_carnot->plan();
  DCHECK(remote_plan);

  DistributedSetSourceGroupGRPCAddressRule set_grpc_address_rule;
  PL_RETURN_IF_ERROR(set_grpc_address_rule.Apply(remote_carnot));

  // Connect the plans.
  for (const auto& [plan, agents] : distributed_plan->plan_to_agent_map()) {
    PL_ASSIGN_OR_RETURN(auto did_connect_plan, AssociateDistributedPlanEdgesRule::ConnectGraphs(
                                                   plan, agents, remote_plan));
    DCHECK(did_connect_plan);
  }

  // TODO(philkuz) make this connect to self without a grpc bridge.
  PL_RETURN_IF_ERROR(
      AssociateDistributedPlanEdgesRule::ConnectGraphs(remote_plan, {remote_node_id}, remote_plan));

  // Expand GRPCSourceGroups in the remote_plan.
  GRPCSourceGroupConversionRule conversion_rule;
  return conversion_rule.Execute(remote_plan).status();
}

StatusOr<std::unique_ptr<DistributedPlan>> DistributedPlanner::Plan(
    const distributedpb::DistributedState& distributed_state, CompilerState*,
    const IR* logical_plan) {
  PL_ASSIGN_OR_RETURN(std::unique_ptr<Coordinator> coordinator,
                      Coordinator::Create(distributed_state));

  PL_ASSIGN_OR_RETURN(std::unique_ptr<DistributedPlan> distributed_plan,
                      coordinator->Coordinate(logical_plan));

  PL_RETURN_IF_ERROR(StitchPlan(distributed_plan.get()));

  AnnotateAbortableSrcsForLimitsRule rule;
  for (IR* agent_plan : distributed_plan->UniquePlans()) {
    rule.Execute(agent_plan);
  }

  return distributed_plan;
}

}  // namespace distributed
}  // namespace planner
}  // namespace carnot
}  // namespace pl
