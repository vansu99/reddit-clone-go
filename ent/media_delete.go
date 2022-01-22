// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ducnguyen96/reddit-clone/ent/media"
	"github.com/ducnguyen96/reddit-clone/ent/predicate"
)

// MediaDelete is the builder for deleting a Media entity.
type MediaDelete struct {
	config
	hooks    []Hook
	mutation *MediaMutation
}

// Where appends a list predicates to the MediaDelete builder.
func (md *MediaDelete) Where(ps ...predicate.Media) *MediaDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MediaDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(md.hooks) == 0 {
		affected, err = md.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			md.mutation = mutation
			affected, err = md.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(md.hooks) - 1; i >= 0; i-- {
			if md.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = md.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, md.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MediaDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MediaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: media.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: media.FieldID,
			},
		},
	}
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, md.driver, _spec)
}

// MediaDeleteOne is the builder for deleting a single Media entity.
type MediaDeleteOne struct {
	md *MediaDelete
}

// Exec executes the deletion query.
func (mdo *MediaDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{media.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MediaDeleteOne) ExecX(ctx context.Context) {
	mdo.md.ExecX(ctx)
}
