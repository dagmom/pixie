package controllers

import (
	"errors"
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"

	storepb "pixielabs.ai/pixielabs/src/vizier/services/metadata/storepb"
	"pixielabs.ai/pixielabs/src/vizier/utils/datastore"
)

const (
	fullResourceUpdatePrefix  = "/fullResourceUpdate"
	topicResourceUpdatePrefix = "/resourceUpdate"
	topicVersionPrefix        = "/topicVersion"
	// The topic for partial resource updates, which are not specific to a particular node.
	unscopedTopic = "unscoped"
)

// MetadataDatastore implements the K8sMetadataStore interface on a given Datastore.
type MetadataDatastore struct {
	ds datastore.MultiGetterSetterDeleterCloser
}

// NewMetadataDatastore wraps the datastore in a metadata store.
func NewMetadataDatastore(ds datastore.MultiGetterSetterDeleterCloser) *MetadataDatastore {
	return &MetadataDatastore{ds}
}

func getFullResourceUpdateKey(version int64) string {
	return path.Join(fullResourceUpdatePrefix, fmt.Sprintf("%020d", version))
}

func getTopicResourceUpdateKey(topic string, version int64) string {
	return path.Join(topicResourceUpdatePrefix, topic, fmt.Sprintf("%020d", version))
}

func getTopicVersionKey(topic string) string {
	return path.Join(topicVersionPrefix, topic)
}

func resourceUpdateKeyToUpdateVersion(key string) (int64, error) {
	splitKey := strings.Split(key, "/")
	if len(splitKey) != 4 {
		return 0, errors.New("Invalid key")
	}
	return strconv.ParseInt(strings.TrimLeft(splitKey[3], "0"), 10, 64)
}

// AddResourceUpdateForTopic stores the given resource with its associated updateVersion for 24h.
func (m *MetadataDatastore) AddResourceUpdateForTopic(updateVersion int64, topic string, resource *storepb.K8SResourceUpdate) error {
	val, err := resource.Marshal()
	if err != nil {
		return err
	}

	return m.ds.SetWithTTL(getTopicResourceUpdateKey(topic, updateVersion), string(val), time.Hour*24)
}

// AddResourceUpdate stores a resource update that is applicable to all topics.
func (m *MetadataDatastore) AddResourceUpdate(updateVersion int64, resource *storepb.K8SResourceUpdate) error {
	return m.AddResourceUpdateForTopic(updateVersion, unscopedTopic, resource)
}

// AddFullResourceUpdate stores full resource update with the given update version.
func (m *MetadataDatastore) AddFullResourceUpdate(updateVersion int64, resource *storepb.K8SResource) error {
	val, err := resource.Marshal()
	if err != nil {
		return err
	}

	return m.ds.SetWithTTL(getFullResourceUpdateKey(updateVersion), string(val), time.Hour*24)
}

// FetchResourceUpdates gets the resource updates from the `from` update version, to the `to`
// update version (exclusive).
func (m *MetadataDatastore) FetchResourceUpdates(topic string, from int64, to int64) ([]*storepb.K8SResourceUpdate, error) {
	// Fetch updates specific to the topic.
	tKeys, tValues, err := m.ds.GetWithRange(getTopicResourceUpdateKey(topic, from), getTopicResourceUpdateKey(topic, to))
	if err != nil {
		return nil, err
	}

	// Fetch updates that are unscoped and shared across all nodes.
	uKeys, uValues, err := m.ds.GetWithRange(getTopicResourceUpdateKey(unscopedTopic, from), getTopicResourceUpdateKey(unscopedTopic, to))
	if err != nil {
		return nil, err
	}

	updates := make([]*storepb.K8SResourceUpdate, 0)

	// Merge the topic-specific and unscoped updates in order.
	tIdx := 0
	uIdx := 0
	for {
		if tIdx == len(tKeys) && uIdx == len(uKeys) {
			break
		}

		var uVersion int64 = -1
		var tVersion int64 = -1
		var err error

		if tIdx < len(tKeys) {
			tVersion, err = resourceUpdateKeyToUpdateVersion(tKeys[tIdx])
			if err != nil { // Malformed key, skip it.
				tIdx++
				continue
			}
		}

		if uIdx < len(uKeys) {
			uVersion, err = resourceUpdateKeyToUpdateVersion(uKeys[uIdx])
			if err != nil { // Malformed key, skip it.
				uIdx++
				continue
			}
		}

		var updateBytes []byte
		if tVersion == -1 || (uVersion != -1 && uVersion < tVersion) {
			updateBytes = uValues[uIdx]
			uIdx++
		} else if uVersion == -1 || (tVersion < uVersion) {
			updateBytes = tValues[tIdx]
			tIdx++
		} else {
			// tVersion == uVersion. This should never happen, since a resource update must either be node-specific or not. In any
			// case, if this ever happens, we should just take the topic-specific update.
			updateBytes = tValues[tIdx]
			tIdx++
			uIdx++
		}

		updatePb := &storepb.K8SResourceUpdate{}
		err = proto.Unmarshal(updateBytes, updatePb)
		if err != nil {
			continue
		}
		updates = append(updates, updatePb)
	}

	return updates, nil
}

// FetchFullResourceUpdates gets the full resource updates from the `from` update version, to the `to`
// update version (exclusive).
func (m *MetadataDatastore) FetchFullResourceUpdates(from int64, to int64) ([]*storepb.K8SResource, error) {
	// Fetch updates specific to the topic.
	_, vals, err := m.ds.GetWithRange(getFullResourceUpdateKey(from), getFullResourceUpdateKey(to))
	if err != nil {
		return nil, err
	}

	updates := make([]*storepb.K8SResource, 0)
	for _, u := range vals {
		updatePb := &storepb.K8SResource{}
		err = proto.Unmarshal(u, updatePb)
		if err != nil {
			continue
		}
		updates = append(updates, updatePb)
	}

	return updates, nil
}

// GetUpdateVersion gets the last update version sent on a topic.
func (m *MetadataDatastore) GetUpdateVersion(topic string) (int64, error) {
	val, err := m.ds.Get(getTopicVersionKey(topic))
	if err != nil {
		return 0, err
	}

	if string(val) == "" {
		return 0, nil
	}

	return strconv.ParseInt(string(val), 10, 64)
}

// SetUpdateVersion sets the last update version sent on a topic.
func (m *MetadataDatastore) SetUpdateVersion(topic string, uv int64) error {
	s := strconv.FormatInt(uv, 10)

	return m.ds.Set(getTopicVersionKey(topic), s)
}