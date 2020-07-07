// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/facebookincubator/ent/examples/o2mrecur/ent/node"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeNode = "Node"
)

// NodeMutation represents an operation that mutate the Nodes
// nodes in the graph.
type NodeMutation struct {
	config
	op              Op
	typ             string
	id              *int
	value           *int
	addvalue        *int
	clearedFields   map[string]struct{}
	parent          *int
	clearedparent   bool
	children        map[int]struct{}
	removedchildren map[int]struct{}
	done            bool
	oldValue        func(context.Context) (*Node, error)
}

var _ ent.Mutation = (*NodeMutation)(nil)

// nodeOption allows to manage the mutation configuration using functional options.
type nodeOption func(*NodeMutation)

// newNodeMutation creates new mutation for $n.Name.
func newNodeMutation(c config, op Op, opts ...nodeOption) *NodeMutation {
	m := &NodeMutation{
		config:        c,
		op:            op,
		typ:           TypeNode,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withNodeID sets the id field of the mutation.
func withNodeID(id int) nodeOption {
	return func(m *NodeMutation) {
		var (
			err   error
			once  sync.Once
			value *Node
		)
		m.oldValue = func(ctx context.Context) (*Node, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Node.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withNode sets the old Node of the mutation.
func withNode(node *Node) nodeOption {
	return func(m *NodeMutation) {
		m.oldValue = func(context.Context) (*Node, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m NodeMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m NodeMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *NodeMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetValue sets the value field.
func (m *NodeMutation) SetValue(i int) {
	m.value = &i
	m.addvalue = nil
}

// Value returns the value value in the mutation.
func (m *NodeMutation) Value() (r int, exists bool) {
	v := m.value
	if v == nil {
		return
	}
	return *v, true
}

// OldValue returns the old value value of the Node.
// If the Node object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *NodeMutation) OldValue(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldValue is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldValue requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldValue: %w", err)
	}
	return oldValue.Value, nil
}

// AddValue adds i to value.
func (m *NodeMutation) AddValue(i int) {
	if m.addvalue != nil {
		*m.addvalue += i
	} else {
		m.addvalue = &i
	}
}

// AddedValue returns the value that was added to the value field in this mutation.
func (m *NodeMutation) AddedValue() (r int, exists bool) {
	v := m.addvalue
	if v == nil {
		return
	}
	return *v, true
}

// ResetValue reset all changes of the "value" field.
func (m *NodeMutation) ResetValue() {
	m.value = nil
	m.addvalue = nil
}

// SetParentID sets the parent edge to Node by id.
func (m *NodeMutation) SetParentID(id int) {
	m.parent = &id
}

// ClearParent clears the parent edge to Node.
func (m *NodeMutation) ClearParent() {
	m.clearedparent = true
}

// ParentCleared returns if the edge parent was cleared.
func (m *NodeMutation) ParentCleared() bool {
	return m.clearedparent
}

// ParentID returns the parent id in the mutation.
func (m *NodeMutation) ParentID() (id int, exists bool) {
	if m.parent != nil {
		return *m.parent, true
	}
	return
}

// ParentIDs returns the parent ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// ParentID instead. It exists only for internal usage by the builders.
func (m *NodeMutation) ParentIDs() (ids []int) {
	if id := m.parent; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetParent reset all changes of the "parent" edge.
func (m *NodeMutation) ResetParent() {
	m.parent = nil
	m.clearedparent = false
}

// AddChildIDs adds the children edge to Node by ids.
func (m *NodeMutation) AddChildIDs(ids ...int) {
	if m.children == nil {
		m.children = make(map[int]struct{})
	}
	for i := range ids {
		m.children[ids[i]] = struct{}{}
	}
}

// RemoveChildIDs removes the children edge to Node by ids.
func (m *NodeMutation) RemoveChildIDs(ids ...int) {
	if m.removedchildren == nil {
		m.removedchildren = make(map[int]struct{})
	}
	for i := range ids {
		m.removedchildren[ids[i]] = struct{}{}
	}
}

// RemovedChildren returns the removed ids of children.
func (m *NodeMutation) RemovedChildrenIDs() (ids []int) {
	for id := range m.removedchildren {
		ids = append(ids, id)
	}
	return
}

// ChildrenIDs returns the children ids in the mutation.
func (m *NodeMutation) ChildrenIDs() (ids []int) {
	for id := range m.children {
		ids = append(ids, id)
	}
	return
}

// ResetChildren reset all changes of the "children" edge.
func (m *NodeMutation) ResetChildren() {
	m.children = nil
	m.removedchildren = nil
}

// Op returns the operation name.
func (m *NodeMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Node).
func (m *NodeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *NodeMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.value != nil {
		fields = append(fields, node.FieldValue)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *NodeMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case node.FieldValue:
		return m.Value()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *NodeMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case node.FieldValue:
		return m.OldValue(ctx)
	}
	return nil, fmt.Errorf("unknown Node field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *NodeMutation) SetField(name string, value ent.Value) error {
	switch name {
	case node.FieldValue:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetValue(v)
		return nil
	}
	return fmt.Errorf("unknown Node field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *NodeMutation) AddedFields() []string {
	var fields []string
	if m.addvalue != nil {
		fields = append(fields, node.FieldValue)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *NodeMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case node.FieldValue:
		return m.AddedValue()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *NodeMutation) AddField(name string, value ent.Value) error {
	switch name {
	case node.FieldValue:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddValue(v)
		return nil
	}
	return fmt.Errorf("unknown Node numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *NodeMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *NodeMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *NodeMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Node nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *NodeMutation) ResetField(name string) error {
	switch name {
	case node.FieldValue:
		m.ResetValue()
		return nil
	}
	return fmt.Errorf("unknown Node field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *NodeMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.parent != nil {
		edges = append(edges, node.EdgeParent)
	}
	if m.children != nil {
		edges = append(edges, node.EdgeChildren)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *NodeMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case node.EdgeParent:
		if id := m.parent; id != nil {
			return []ent.Value{*id}
		}
	case node.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.children))
		for id := range m.children {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *NodeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedchildren != nil {
		edges = append(edges, node.EdgeChildren)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *NodeMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case node.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.removedchildren))
		for id := range m.removedchildren {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *NodeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedparent {
		edges = append(edges, node.EdgeParent)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *NodeMutation) EdgeCleared(name string) bool {
	switch name {
	case node.EdgeParent:
		return m.clearedparent
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *NodeMutation) ClearEdge(name string) error {
	switch name {
	case node.EdgeParent:
		m.ClearParent()
		return nil
	}
	return fmt.Errorf("unknown Node unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *NodeMutation) ResetEdge(name string) error {
	switch name {
	case node.EdgeParent:
		m.ResetParent()
		return nil
	case node.EdgeChildren:
		m.ResetChildren()
		return nil
	}
	return fmt.Errorf("unknown Node edge %s", name)
}
