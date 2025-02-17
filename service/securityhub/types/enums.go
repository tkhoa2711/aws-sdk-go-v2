// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdminStatus string

// Enum values for AdminStatus
const (
	AdminStatusEnabled           AdminStatus = "ENABLED"
	AdminStatusDisableInProgress AdminStatus = "DISABLE_IN_PROGRESS"
)

// Values returns all known values for AdminStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AdminStatus) Values() []AdminStatus {
	return []AdminStatus{
		"ENABLED",
		"DISABLE_IN_PROGRESS",
	}
}

type AwsIamAccessKeyStatus string

// Enum values for AwsIamAccessKeyStatus
const (
	AwsIamAccessKeyStatusActive   AwsIamAccessKeyStatus = "Active"
	AwsIamAccessKeyStatusInactive AwsIamAccessKeyStatus = "Inactive"
)

// Values returns all known values for AwsIamAccessKeyStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AwsIamAccessKeyStatus) Values() []AwsIamAccessKeyStatus {
	return []AwsIamAccessKeyStatus{
		"Active",
		"Inactive",
	}
}

type AwsS3BucketNotificationConfigurationS3KeyFilterRuleName string

// Enum values for AwsS3BucketNotificationConfigurationS3KeyFilterRuleName
const (
	AwsS3BucketNotificationConfigurationS3KeyFilterRuleNamePrefix AwsS3BucketNotificationConfigurationS3KeyFilterRuleName = "Prefix"
	AwsS3BucketNotificationConfigurationS3KeyFilterRuleNameSuffix AwsS3BucketNotificationConfigurationS3KeyFilterRuleName = "Suffix"
)

// Values returns all known values for
// AwsS3BucketNotificationConfigurationS3KeyFilterRuleName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AwsS3BucketNotificationConfigurationS3KeyFilterRuleName) Values() []AwsS3BucketNotificationConfigurationS3KeyFilterRuleName {
	return []AwsS3BucketNotificationConfigurationS3KeyFilterRuleName{
		"Prefix",
		"Suffix",
	}
}

type ComplianceStatus string

// Enum values for ComplianceStatus
const (
	ComplianceStatusPassed       ComplianceStatus = "PASSED"
	ComplianceStatusWarning      ComplianceStatus = "WARNING"
	ComplianceStatusFailed       ComplianceStatus = "FAILED"
	ComplianceStatusNotAvailable ComplianceStatus = "NOT_AVAILABLE"
)

// Values returns all known values for ComplianceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComplianceStatus) Values() []ComplianceStatus {
	return []ComplianceStatus{
		"PASSED",
		"WARNING",
		"FAILED",
		"NOT_AVAILABLE",
	}
}

type ControlStatus string

// Enum values for ControlStatus
const (
	ControlStatusEnabled  ControlStatus = "ENABLED"
	ControlStatusDisabled ControlStatus = "DISABLED"
)

// Values returns all known values for ControlStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ControlStatus) Values() []ControlStatus {
	return []ControlStatus{
		"ENABLED",
		"DISABLED",
	}
}

type DateRangeUnit string

// Enum values for DateRangeUnit
const (
	DateRangeUnitDays DateRangeUnit = "DAYS"
)

// Values returns all known values for DateRangeUnit. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DateRangeUnit) Values() []DateRangeUnit {
	return []DateRangeUnit{
		"DAYS",
	}
}

type IntegrationType string

// Enum values for IntegrationType
const (
	IntegrationTypeSendFindingsToSecurityHub      IntegrationType = "SEND_FINDINGS_TO_SECURITY_HUB"
	IntegrationTypeReceiveFindingsFromSecurityHub IntegrationType = "RECEIVE_FINDINGS_FROM_SECURITY_HUB"
	IntegrationTypeUpdateFindingsInSecurityHub    IntegrationType = "UPDATE_FINDINGS_IN_SECURITY_HUB"
)

// Values returns all known values for IntegrationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IntegrationType) Values() []IntegrationType {
	return []IntegrationType{
		"SEND_FINDINGS_TO_SECURITY_HUB",
		"RECEIVE_FINDINGS_FROM_SECURITY_HUB",
		"UPDATE_FINDINGS_IN_SECURITY_HUB",
	}
}

type MalwareState string

// Enum values for MalwareState
const (
	MalwareStateObserved      MalwareState = "OBSERVED"
	MalwareStateRemovalFailed MalwareState = "REMOVAL_FAILED"
	MalwareStateRemoved       MalwareState = "REMOVED"
)

// Values returns all known values for MalwareState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MalwareState) Values() []MalwareState {
	return []MalwareState{
		"OBSERVED",
		"REMOVAL_FAILED",
		"REMOVED",
	}
}

type MalwareType string

// Enum values for MalwareType
const (
	MalwareTypeAdware              MalwareType = "ADWARE"
	MalwareTypeBlendedThreat       MalwareType = "BLENDED_THREAT"
	MalwareTypeBotnetAgent         MalwareType = "BOTNET_AGENT"
	MalwareTypeCoinMiner           MalwareType = "COIN_MINER"
	MalwareTypeExploitKit          MalwareType = "EXPLOIT_KIT"
	MalwareTypeKeylogger           MalwareType = "KEYLOGGER"
	MalwareTypeMacro               MalwareType = "MACRO"
	MalwareTypePotentiallyUnwanted MalwareType = "POTENTIALLY_UNWANTED"
	MalwareTypeSpyware             MalwareType = "SPYWARE"
	MalwareTypeRansomware          MalwareType = "RANSOMWARE"
	MalwareTypeRemoteAccess        MalwareType = "REMOTE_ACCESS"
	MalwareTypeRootkit             MalwareType = "ROOTKIT"
	MalwareTypeTrojan              MalwareType = "TROJAN"
	MalwareTypeVirus               MalwareType = "VIRUS"
	MalwareTypeWorm                MalwareType = "WORM"
)

// Values returns all known values for MalwareType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MalwareType) Values() []MalwareType {
	return []MalwareType{
		"ADWARE",
		"BLENDED_THREAT",
		"BOTNET_AGENT",
		"COIN_MINER",
		"EXPLOIT_KIT",
		"KEYLOGGER",
		"MACRO",
		"POTENTIALLY_UNWANTED",
		"SPYWARE",
		"RANSOMWARE",
		"REMOTE_ACCESS",
		"ROOTKIT",
		"TROJAN",
		"VIRUS",
		"WORM",
	}
}

type MapFilterComparison string

// Enum values for MapFilterComparison
const (
	MapFilterComparisonEquals    MapFilterComparison = "EQUALS"
	MapFilterComparisonNotEquals MapFilterComparison = "NOT_EQUALS"
)

// Values returns all known values for MapFilterComparison. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MapFilterComparison) Values() []MapFilterComparison {
	return []MapFilterComparison{
		"EQUALS",
		"NOT_EQUALS",
	}
}

type NetworkDirection string

// Enum values for NetworkDirection
const (
	NetworkDirectionIn  NetworkDirection = "IN"
	NetworkDirectionOut NetworkDirection = "OUT"
)

// Values returns all known values for NetworkDirection. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NetworkDirection) Values() []NetworkDirection {
	return []NetworkDirection{
		"IN",
		"OUT",
	}
}

type Partition string

// Enum values for Partition
const (
	PartitionAws      Partition = "aws"
	PartitionAwsCn    Partition = "aws-cn"
	PartitionAwsUsGov Partition = "aws-us-gov"
)

// Values returns all known values for Partition. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Partition) Values() []Partition {
	return []Partition{
		"aws",
		"aws-cn",
		"aws-us-gov",
	}
}

type RecordState string

// Enum values for RecordState
const (
	RecordStateActive   RecordState = "ACTIVE"
	RecordStateArchived RecordState = "ARCHIVED"
)

// Values returns all known values for RecordState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RecordState) Values() []RecordState {
	return []RecordState{
		"ACTIVE",
		"ARCHIVED",
	}
}

type SeverityLabel string

// Enum values for SeverityLabel
const (
	SeverityLabelInformational SeverityLabel = "INFORMATIONAL"
	SeverityLabelLow           SeverityLabel = "LOW"
	SeverityLabelMedium        SeverityLabel = "MEDIUM"
	SeverityLabelHigh          SeverityLabel = "HIGH"
	SeverityLabelCritical      SeverityLabel = "CRITICAL"
)

// Values returns all known values for SeverityLabel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SeverityLabel) Values() []SeverityLabel {
	return []SeverityLabel{
		"INFORMATIONAL",
		"LOW",
		"MEDIUM",
		"HIGH",
		"CRITICAL",
	}
}

type SeverityRating string

// Enum values for SeverityRating
const (
	SeverityRatingLow      SeverityRating = "LOW"
	SeverityRatingMedium   SeverityRating = "MEDIUM"
	SeverityRatingHigh     SeverityRating = "HIGH"
	SeverityRatingCritical SeverityRating = "CRITICAL"
)

// Values returns all known values for SeverityRating. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SeverityRating) Values() []SeverityRating {
	return []SeverityRating{
		"LOW",
		"MEDIUM",
		"HIGH",
		"CRITICAL",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "asc"
	SortOrderDescending SortOrder = "desc"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"asc",
		"desc",
	}
}

type StandardsStatus string

// Enum values for StandardsStatus
const (
	StandardsStatusPending    StandardsStatus = "PENDING"
	StandardsStatusReady      StandardsStatus = "READY"
	StandardsStatusFailed     StandardsStatus = "FAILED"
	StandardsStatusDeleting   StandardsStatus = "DELETING"
	StandardsStatusIncomplete StandardsStatus = "INCOMPLETE"
)

// Values returns all known values for StandardsStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StandardsStatus) Values() []StandardsStatus {
	return []StandardsStatus{
		"PENDING",
		"READY",
		"FAILED",
		"DELETING",
		"INCOMPLETE",
	}
}

type StringFilterComparison string

// Enum values for StringFilterComparison
const (
	StringFilterComparisonEquals          StringFilterComparison = "EQUALS"
	StringFilterComparisonPrefix          StringFilterComparison = "PREFIX"
	StringFilterComparisonNotEquals       StringFilterComparison = "NOT_EQUALS"
	StringFilterComparisonPrefixNotEquals StringFilterComparison = "PREFIX_NOT_EQUALS"
)

// Values returns all known values for StringFilterComparison. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StringFilterComparison) Values() []StringFilterComparison {
	return []StringFilterComparison{
		"EQUALS",
		"PREFIX",
		"NOT_EQUALS",
		"PREFIX_NOT_EQUALS",
	}
}

type ThreatIntelIndicatorCategory string

// Enum values for ThreatIntelIndicatorCategory
const (
	ThreatIntelIndicatorCategoryBackdoor          ThreatIntelIndicatorCategory = "BACKDOOR"
	ThreatIntelIndicatorCategoryCardStealer       ThreatIntelIndicatorCategory = "CARD_STEALER"
	ThreatIntelIndicatorCategoryCommandAndControl ThreatIntelIndicatorCategory = "COMMAND_AND_CONTROL"
	ThreatIntelIndicatorCategoryDropSite          ThreatIntelIndicatorCategory = "DROP_SITE"
	ThreatIntelIndicatorCategoryExploitSite       ThreatIntelIndicatorCategory = "EXPLOIT_SITE"
	ThreatIntelIndicatorCategoryKeylogger         ThreatIntelIndicatorCategory = "KEYLOGGER"
)

// Values returns all known values for ThreatIntelIndicatorCategory. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ThreatIntelIndicatorCategory) Values() []ThreatIntelIndicatorCategory {
	return []ThreatIntelIndicatorCategory{
		"BACKDOOR",
		"CARD_STEALER",
		"COMMAND_AND_CONTROL",
		"DROP_SITE",
		"EXPLOIT_SITE",
		"KEYLOGGER",
	}
}

type ThreatIntelIndicatorType string

// Enum values for ThreatIntelIndicatorType
const (
	ThreatIntelIndicatorTypeDomain       ThreatIntelIndicatorType = "DOMAIN"
	ThreatIntelIndicatorTypeEmailAddress ThreatIntelIndicatorType = "EMAIL_ADDRESS"
	ThreatIntelIndicatorTypeHashMd5      ThreatIntelIndicatorType = "HASH_MD5"
	ThreatIntelIndicatorTypeHashSha1     ThreatIntelIndicatorType = "HASH_SHA1"
	ThreatIntelIndicatorTypeHashSha256   ThreatIntelIndicatorType = "HASH_SHA256"
	ThreatIntelIndicatorTypeHashSha512   ThreatIntelIndicatorType = "HASH_SHA512"
	ThreatIntelIndicatorTypeIpv4Address  ThreatIntelIndicatorType = "IPV4_ADDRESS"
	ThreatIntelIndicatorTypeIpv6Address  ThreatIntelIndicatorType = "IPV6_ADDRESS"
	ThreatIntelIndicatorTypeMutex        ThreatIntelIndicatorType = "MUTEX"
	ThreatIntelIndicatorTypeProcess      ThreatIntelIndicatorType = "PROCESS"
	ThreatIntelIndicatorTypeUrl          ThreatIntelIndicatorType = "URL"
)

// Values returns all known values for ThreatIntelIndicatorType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ThreatIntelIndicatorType) Values() []ThreatIntelIndicatorType {
	return []ThreatIntelIndicatorType{
		"DOMAIN",
		"EMAIL_ADDRESS",
		"HASH_MD5",
		"HASH_SHA1",
		"HASH_SHA256",
		"HASH_SHA512",
		"IPV4_ADDRESS",
		"IPV6_ADDRESS",
		"MUTEX",
		"PROCESS",
		"URL",
	}
}

type VerificationState string

// Enum values for VerificationState
const (
	VerificationStateUnknown        VerificationState = "UNKNOWN"
	VerificationStateTruePositive   VerificationState = "TRUE_POSITIVE"
	VerificationStateFalsePositive  VerificationState = "FALSE_POSITIVE"
	VerificationStateBenignPositive VerificationState = "BENIGN_POSITIVE"
)

// Values returns all known values for VerificationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VerificationState) Values() []VerificationState {
	return []VerificationState{
		"UNKNOWN",
		"TRUE_POSITIVE",
		"FALSE_POSITIVE",
		"BENIGN_POSITIVE",
	}
}

type WorkflowState string

// Enum values for WorkflowState
const (
	WorkflowStateNew        WorkflowState = "NEW"
	WorkflowStateAssigned   WorkflowState = "ASSIGNED"
	WorkflowStateInProgress WorkflowState = "IN_PROGRESS"
	WorkflowStateDeferred   WorkflowState = "DEFERRED"
	WorkflowStateResolved   WorkflowState = "RESOLVED"
)

// Values returns all known values for WorkflowState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowState) Values() []WorkflowState {
	return []WorkflowState{
		"NEW",
		"ASSIGNED",
		"IN_PROGRESS",
		"DEFERRED",
		"RESOLVED",
	}
}

type WorkflowStatus string

// Enum values for WorkflowStatus
const (
	WorkflowStatusNew        WorkflowStatus = "NEW"
	WorkflowStatusNotified   WorkflowStatus = "NOTIFIED"
	WorkflowStatusResolved   WorkflowStatus = "RESOLVED"
	WorkflowStatusSuppressed WorkflowStatus = "SUPPRESSED"
)

// Values returns all known values for WorkflowStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowStatus) Values() []WorkflowStatus {
	return []WorkflowStatus{
		"NEW",
		"NOTIFIED",
		"RESOLVED",
		"SUPPRESSED",
	}
}
