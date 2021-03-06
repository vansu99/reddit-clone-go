// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/ducnguyen96/reddit-clone/ent/action"
	"github.com/ducnguyen96/reddit-clone/ent/comment"
	"github.com/ducnguyen96/reddit-clone/ent/community"
	"github.com/ducnguyen96/reddit-clone/ent/media"
	"github.com/ducnguyen96/reddit-clone/ent/post"
	"github.com/ducnguyen96/reddit-clone/ent/schema"
	"github.com/ducnguyen96/reddit-clone/ent/tag"
	"github.com/ducnguyen96/reddit-clone/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	actionMixin := schema.Action{}.Mixin()
	actionMixinHooks0 := actionMixin[0].Hooks()
	action.Hooks[0] = actionMixinHooks0[0]
	actionMixinFields0 := actionMixin[0].Fields()
	_ = actionMixinFields0
	actionFields := schema.Action{}.Fields()
	_ = actionFields
	// actionDescCreatedAt is the schema descriptor for created_at field.
	actionDescCreatedAt := actionMixinFields0[1].Descriptor()
	// action.DefaultCreatedAt holds the default value on creation for the created_at field.
	action.DefaultCreatedAt = actionDescCreatedAt.Default.(func() time.Time)
	// actionDescUpdatedAt is the schema descriptor for updated_at field.
	actionDescUpdatedAt := actionMixinFields0[2].Descriptor()
	// action.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	action.DefaultUpdatedAt = actionDescUpdatedAt.Default.(func() time.Time)
	// action.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	action.UpdateDefaultUpdatedAt = actionDescUpdatedAt.UpdateDefault.(func() time.Time)
	commentMixin := schema.Comment{}.Mixin()
	commentMixinHooks0 := commentMixin[0].Hooks()
	comment.Hooks[0] = commentMixinHooks0[0]
	commentMixinFields0 := commentMixin[0].Fields()
	_ = commentMixinFields0
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentMixinFields0[1].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	// commentDescUpdatedAt is the schema descriptor for updated_at field.
	commentDescUpdatedAt := commentMixinFields0[2].Descriptor()
	// comment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	comment.DefaultUpdatedAt = commentDescUpdatedAt.Default.(func() time.Time)
	// comment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	comment.UpdateDefaultUpdatedAt = commentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// commentDescUpVotes is the schema descriptor for up_votes field.
	commentDescUpVotes := commentFields[2].Descriptor()
	// comment.DefaultUpVotes holds the default value on creation for the up_votes field.
	comment.DefaultUpVotes = commentDescUpVotes.Default.(int)
	// commentDescDownVotes is the schema descriptor for down_votes field.
	commentDescDownVotes := commentFields[3].Descriptor()
	// comment.DefaultDownVotes holds the default value on creation for the down_votes field.
	comment.DefaultDownVotes = commentDescDownVotes.Default.(int)
	communityMixin := schema.Community{}.Mixin()
	communityMixinHooks0 := communityMixin[0].Hooks()
	community.Hooks[0] = communityMixinHooks0[0]
	communityMixinFields0 := communityMixin[0].Fields()
	_ = communityMixinFields0
	communityFields := schema.Community{}.Fields()
	_ = communityFields
	// communityDescCreatedAt is the schema descriptor for created_at field.
	communityDescCreatedAt := communityMixinFields0[1].Descriptor()
	// community.DefaultCreatedAt holds the default value on creation for the created_at field.
	community.DefaultCreatedAt = communityDescCreatedAt.Default.(func() time.Time)
	// communityDescUpdatedAt is the schema descriptor for updated_at field.
	communityDescUpdatedAt := communityMixinFields0[2].Descriptor()
	// community.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	community.DefaultUpdatedAt = communityDescUpdatedAt.Default.(func() time.Time)
	// community.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	community.UpdateDefaultUpdatedAt = communityDescUpdatedAt.UpdateDefault.(func() time.Time)
	// communityDescName is the schema descriptor for name field.
	communityDescName := communityFields[0].Descriptor()
	// community.NameValidator is a validator for the "name" field. It is called by the builders before save.
	community.NameValidator = func() func(string) error {
		validators := communityDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// communityDescSlug is the schema descriptor for slug field.
	communityDescSlug := communityFields[1].Descriptor()
	// community.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	community.SlugValidator = func() func(string) error {
		validators := communityDescSlug.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(slug string) error {
			for _, fn := range fns {
				if err := fn(slug); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// communityDescIsAdult is the schema descriptor for is_adult field.
	communityDescIsAdult := communityFields[3].Descriptor()
	// community.DefaultIsAdult holds the default value on creation for the is_adult field.
	community.DefaultIsAdult = communityDescIsAdult.Default.(bool)
	mediaMixin := schema.Media{}.Mixin()
	mediaMixinHooks0 := mediaMixin[0].Hooks()
	media.Hooks[0] = mediaMixinHooks0[0]
	mediaMixinFields0 := mediaMixin[0].Fields()
	_ = mediaMixinFields0
	mediaFields := schema.Media{}.Fields()
	_ = mediaFields
	// mediaDescCreatedAt is the schema descriptor for created_at field.
	mediaDescCreatedAt := mediaMixinFields0[1].Descriptor()
	// media.DefaultCreatedAt holds the default value on creation for the created_at field.
	media.DefaultCreatedAt = mediaDescCreatedAt.Default.(func() time.Time)
	// mediaDescUpdatedAt is the schema descriptor for updated_at field.
	mediaDescUpdatedAt := mediaMixinFields0[2].Descriptor()
	// media.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	media.DefaultUpdatedAt = mediaDescUpdatedAt.Default.(func() time.Time)
	// media.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	media.UpdateDefaultUpdatedAt = mediaDescUpdatedAt.UpdateDefault.(func() time.Time)
	postMixin := schema.Post{}.Mixin()
	postMixinHooks0 := postMixin[0].Hooks()
	post.Hooks[0] = postMixinHooks0[0]
	postMixinFields0 := postMixin[0].Fields()
	_ = postMixinFields0
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postMixinFields0[1].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for updated_at field.
	postDescUpdatedAt := postMixinFields0[2].Descriptor()
	// post.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)
	// post.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[0].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = func() func(string) error {
		validators := postDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// postDescSlug is the schema descriptor for slug field.
	postDescSlug := postFields[1].Descriptor()
	// post.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	post.SlugValidator = func() func(string) error {
		validators := postDescSlug.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(slug string) error {
			for _, fn := range fns {
				if err := fn(slug); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// postDescUpVotes is the schema descriptor for up_votes field.
	postDescUpVotes := postFields[5].Descriptor()
	// post.DefaultUpVotes holds the default value on creation for the up_votes field.
	post.DefaultUpVotes = postDescUpVotes.Default.(int)
	// postDescDownVotes is the schema descriptor for down_votes field.
	postDescDownVotes := postFields[6].Descriptor()
	// post.DefaultDownVotes holds the default value on creation for the down_votes field.
	post.DefaultDownVotes = postDescDownVotes.Default.(int)
	tagMixin := schema.Tag{}.Mixin()
	tagMixinHooks0 := tagMixin[0].Hooks()
	tag.Hooks[0] = tagMixinHooks0[0]
	tagMixinFields0 := tagMixin[0].Fields()
	_ = tagMixinFields0
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagMixinFields0[1].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	// tagDescUpdatedAt is the schema descriptor for updated_at field.
	tagDescUpdatedAt := tagMixinFields0[2].Descriptor()
	// tag.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tag.DefaultUpdatedAt = tagDescUpdatedAt.Default.(func() time.Time)
	// tag.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	tag.UpdateDefaultUpdatedAt = tagDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tagDescValue is the schema descriptor for value field.
	tagDescValue := tagFields[0].Descriptor()
	// tag.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	tag.ValueValidator = func() func(string) error {
		validators := tagDescValue.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(value string) error {
			for _, fn := range fns {
				if err := fn(value); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}

const (
	Version = "v0.9.1"                                          // Version of ent codegen.
	Sum     = "h1:IG8andyeD79GG24U8Q+1Y45hQXj6gY5evSBcva5gtBk=" // Sum of ent codegen.
)
