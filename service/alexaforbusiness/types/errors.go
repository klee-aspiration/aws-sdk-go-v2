// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The resource being created already exists.
type AlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyExistsException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "AlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *AlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There is a concurrent modification of resources.
type ConcurrentModificationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ConcurrentModificationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request failed because this device is no longer registered and therefore no
// longer managed by this account.
type DeviceNotRegisteredException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DeviceNotRegisteredException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DeviceNotRegisteredException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DeviceNotRegisteredException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "DeviceNotRegisteredException"
	}
	return *e.ErrorCodeOverride
}
func (e *DeviceNotRegisteredException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Certificate Authority can't issue or revoke a certificate.
type InvalidCertificateAuthorityException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidCertificateAuthorityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCertificateAuthorityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCertificateAuthorityException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidCertificateAuthorityException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidCertificateAuthorityException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The device is in an invalid state.
type InvalidDeviceException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidDeviceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDeviceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDeviceException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidDeviceException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidDeviceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A password in SecretsManager is in an invalid state.
type InvalidSecretsManagerResourceException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidSecretsManagerResourceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSecretsManagerResourceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSecretsManagerResourceException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidSecretsManagerResourceException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSecretsManagerResourceException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The service linked role is locked for deletion.
type InvalidServiceLinkedRoleStateException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidServiceLinkedRoleStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidServiceLinkedRoleStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidServiceLinkedRoleStateException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidServiceLinkedRoleStateException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidServiceLinkedRoleStateException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The attempt to update a user is invalid due to the user's current status.
type InvalidUserStatusException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidUserStatusException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidUserStatusException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidUserStatusException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidUserStatusException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidUserStatusException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You are performing an action that would put you beyond your account's limits.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The name sent in the request is already in use.
type NameInUseException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NameInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NameInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NameInUseException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "NameInUseException"
	}
	return *e.ErrorCodeOverride
}
func (e *NameInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource is not found.
type NotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "NotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Another resource is associated with the resource in the request.
type ResourceAssociatedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceAssociatedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAssociatedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAssociatedException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceAssociatedException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceAssociatedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource in the request is already in use.
type ResourceInUseException struct {
	Message *string

	ErrorCodeOverride *string

	ClientRequestToken *string

	noSmithyDocumentSerde
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceInUseException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The skill must be linked to a third-party account.
type SkillNotLinkedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SkillNotLinkedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SkillNotLinkedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SkillNotLinkedException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "SkillNotLinkedException"
	}
	return *e.ErrorCodeOverride
}
func (e *SkillNotLinkedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The caller has no permissions to operate on the resource involved in the API
// call.
type UnauthorizedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UnauthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "UnauthorizedException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnauthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
