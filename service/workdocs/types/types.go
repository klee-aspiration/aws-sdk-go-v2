// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Describes the activity information.
type Activity struct {

	// Metadata of the commenting activity. This is an optional field and is filled for
	// commenting activities.
	CommentMetadata *CommentMetadata

	// The user who performed the action.
	Initiator *UserMetadata

	// Indicates whether an activity is indirect or direct. An indirect activity
	// results from a direct activity performed on a parent resource. For example,
	// sharing a parent folder (the direct activity) shares all of the subfolders and
	// documents within the parent folder (the indirect activity).
	IsIndirectActivity bool

	// The ID of the organization.
	OrganizationId *string

	// The original parent of the resource. This is an optional field and is filled for
	// move activities.
	OriginalParent *ResourceMetadata

	// The list of users or groups impacted by this action. This is an optional field
	// and is filled for the following sharing activities: DOCUMENT_SHARED,
	// DOCUMENT_SHARED, DOCUMENT_UNSHARED, FOLDER_SHARED, FOLDER_UNSHARED.
	Participants *Participants

	// The metadata of the resource involved in the user action.
	ResourceMetadata *ResourceMetadata

	// The timestamp when the action was performed.
	TimeStamp *time.Time

	// The activity type.
	Type ActivityType

	noSmithyDocumentSerde
}

// Describes a comment.
type Comment struct {

	// The ID of the comment.
	//
	// This member is required.
	CommentId *string

	// The details of the user who made the comment.
	Contributor *User

	// The time that the comment was created.
	CreatedTimestamp *time.Time

	// The ID of the parent comment.
	ParentId *string

	// If the comment is a reply to another user's comment, this field contains the
	// user ID of the user being replied to.
	RecipientId *string

	// The status of the comment.
	Status CommentStatusType

	// The text of the comment.
	Text *string

	// The ID of the root comment in the thread.
	ThreadId *string

	// The visibility of the comment. Options are either PRIVATE, where the comment is
	// visible only to the comment author and document owner and co-owners, or PUBLIC,
	// where the comment is visible to document owners, co-owners, and contributors.
	Visibility CommentVisibilityType

	noSmithyDocumentSerde
}

// Describes the metadata of a comment.
type CommentMetadata struct {

	// The ID of the comment.
	CommentId *string

	// The status of the comment.
	CommentStatus CommentStatusType

	// The user who made the comment.
	Contributor *User

	// The timestamp that the comment was created.
	CreatedTimestamp *time.Time

	// The ID of the user being replied to.
	RecipientId *string

	noSmithyDocumentSerde
}

// Describes the document.
type DocumentMetadata struct {

	// The time when the document was created.
	CreatedTimestamp *time.Time

	// The ID of the creator.
	CreatorId *string

	// The ID of the document.
	Id *string

	// List of labels on the document.
	Labels []string

	// The latest version of the document.
	LatestVersionMetadata *DocumentVersionMetadata

	// The time when the document was updated.
	ModifiedTimestamp *time.Time

	// The ID of the parent folder.
	ParentFolderId *string

	// The resource state.
	ResourceState ResourceStateType

	noSmithyDocumentSerde
}

// Describes a version of a document.
type DocumentVersionMetadata struct {

	// The timestamp when the content of the document was originally created.
	ContentCreatedTimestamp *time.Time

	// The timestamp when the content of the document was modified.
	ContentModifiedTimestamp *time.Time

	// The content type of the document.
	ContentType *string

	// The timestamp when the document was first uploaded.
	CreatedTimestamp *time.Time

	// The ID of the creator.
	CreatorId *string

	// The ID of the version.
	Id *string

	// The timestamp when the document was last uploaded.
	ModifiedTimestamp *time.Time

	// The name of the version.
	Name *string

	// The signature of the document.
	Signature *string

	// The size of the document, in bytes.
	Size *int64

	// The source of the document.
	Source map[string]string

	// The status of the document.
	Status DocumentStatusType

	// The thumbnail of the document.
	Thumbnail map[string]string

	noSmithyDocumentSerde
}

// Describes a folder.
type FolderMetadata struct {

	// The time when the folder was created.
	CreatedTimestamp *time.Time

	// The ID of the creator.
	CreatorId *string

	// The ID of the folder.
	Id *string

	// List of labels on the folder.
	Labels []string

	// The size of the latest version of the folder metadata.
	LatestVersionSize *int64

	// The time when the folder was updated.
	ModifiedTimestamp *time.Time

	// The name of the folder.
	Name *string

	// The ID of the parent folder.
	ParentFolderId *string

	// The resource state of the folder.
	ResourceState ResourceStateType

	// The unique identifier created from the subfolders and documents of the folder.
	Signature *string

	// The size of the folder metadata.
	Size *int64

	noSmithyDocumentSerde
}

// Describes the metadata of a user group.
type GroupMetadata struct {

	// The ID of the user group.
	Id *string

	// The name of the group.
	Name *string

	noSmithyDocumentSerde
}

// Set of options which defines notification preferences of given action.
type NotificationOptions struct {

	// Text value to be included in the email body.
	EmailMessage *string

	// Boolean value to indicate an email notification should be sent to the
	// recipients.
	SendEmail bool

	noSmithyDocumentSerde
}

// Describes the users or user groups.
type Participants struct {

	// The list of user groups.
	Groups []GroupMetadata

	// The list of users.
	Users []UserMetadata

	noSmithyDocumentSerde
}

// Describes the permissions.
type PermissionInfo struct {

	// The role of the user.
	Role RoleType

	// The type of permissions.
	Type RolePermissionType

	noSmithyDocumentSerde
}

// Describes a resource.
type Principal struct {

	// The ID of the resource.
	Id *string

	// The permission information for the resource.
	Roles []PermissionInfo

	// The type of resource.
	Type PrincipalType

	noSmithyDocumentSerde
}

// Describes the metadata of a resource.
type ResourceMetadata struct {

	// The ID of the resource.
	Id *string

	// The name of the resource.
	Name *string

	// The original name of the resource before a rename operation.
	OriginalName *string

	// The owner of the resource.
	Owner *UserMetadata

	// The parent ID of the resource before a rename operation.
	ParentId *string

	// The type of resource.
	Type ResourceType

	// The version ID of the resource. This is an optional field and is filled for
	// action on document version.
	VersionId *string

	noSmithyDocumentSerde
}

// Describes the path information of a resource.
type ResourcePath struct {

	// The components of the resource path.
	Components []ResourcePathComponent

	noSmithyDocumentSerde
}

// Describes the resource path.
type ResourcePathComponent struct {

	// The ID of the resource path.
	Id *string

	// The name of the resource path.
	Name *string

	noSmithyDocumentSerde
}

// Describes the recipient type and ID, if available.
type SharePrincipal struct {

	// The ID of the recipient.
	//
	// This member is required.
	Id *string

	// The role of the recipient.
	//
	// This member is required.
	Role RoleType

	// The type of the recipient.
	//
	// This member is required.
	Type PrincipalType

	noSmithyDocumentSerde
}

// Describes the share results of a resource.
type ShareResult struct {

	// The ID of the invited user.
	InviteePrincipalId *string

	// The ID of the principal.
	PrincipalId *string

	// The role.
	Role RoleType

	// The ID of the resource that was shared.
	ShareId *string

	// The status.
	Status ShareStatusType

	// The status message.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Describes the storage for a user.
type StorageRuleType struct {

	// The amount of storage allocated, in bytes.
	StorageAllocatedInBytes *int64

	// The type of storage.
	StorageType StorageType

	noSmithyDocumentSerde
}

// Describes a subscription.
type Subscription struct {

	// The endpoint of the subscription.
	EndPoint *string

	// The protocol of the subscription.
	Protocol SubscriptionProtocolType

	// The ID of the subscription.
	SubscriptionId *string

	noSmithyDocumentSerde
}

// Describes the upload.
type UploadMetadata struct {

	// The signed headers.
	SignedHeaders map[string]string

	// The URL of the upload.
	UploadUrl *string

	noSmithyDocumentSerde
}

// Describes a user.
type User struct {

	// The time when the user was created.
	CreatedTimestamp *time.Time

	// The email address of the user.
	EmailAddress *string

	// The given name of the user.
	GivenName *string

	// The ID of the user.
	Id *string

	// The locale of the user.
	Locale LocaleType

	// The time when the user was modified.
	ModifiedTimestamp *time.Time

	// The ID of the organization.
	OrganizationId *string

	// The ID of the recycle bin folder.
	RecycleBinFolderId *string

	// The ID of the root folder.
	RootFolderId *string

	// The status of the user.
	Status UserStatusType

	// The storage for the user.
	Storage *UserStorageMetadata

	// The surname of the user.
	Surname *string

	// The time zone ID of the user.
	TimeZoneId *string

	// The type of user.
	Type UserType

	// The login name of the user.
	Username *string

	noSmithyDocumentSerde
}

// Describes the metadata of the user.
type UserMetadata struct {

	// The email address of the user.
	EmailAddress *string

	// The given name of the user before a rename operation.
	GivenName *string

	// The ID of the user.
	Id *string

	// The surname of the user.
	Surname *string

	// The name of the user.
	Username *string

	noSmithyDocumentSerde
}

// Describes the storage for a user.
type UserStorageMetadata struct {

	// The storage for a user.
	StorageRule *StorageRuleType

	// The amount of storage used, in bytes.
	StorageUtilizedInBytes *int64

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
