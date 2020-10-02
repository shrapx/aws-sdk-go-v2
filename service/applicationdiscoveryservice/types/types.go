// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Information about agents or connectors that were instructed to start collecting
// data. Information includes the agent/connector ID, a description of the
// operation, and whether the agent/connector configuration was updated.
type AgentConfigurationStatus struct {

	// The agent/connector ID.
	AgentId *string

	// A description of the operation performed.
	Description *string

	// Information about the status of the StartDataCollection and StopDataCollection
	// operations. The system has recorded the data collection operation. The
	// agent/connector receives this command the next time it polls for a new command.
	OperationSucceeded *bool
}

// Information about agents or connectors associated with the user’s AWS account.
// Information includes agent or connector IDs, IP addresses, media access control
// (MAC) addresses, agent or connector health, hostname where the agent or
// connector resides, and agent version for each agent.
type AgentInfo struct {

	// The agent or connector ID.
	AgentId *string

	// Network details about the host where the agent or connector resides.
	AgentNetworkInfoList []*AgentNetworkInfo

	// Type of agent.
	AgentType *string

	// Status of the collection process for an agent or connector.
	CollectionStatus *string

	// The ID of the connector.
	ConnectorId *string

	// The health of the agent or connector.
	Health AgentStatus

	// The name of the host where the agent or connector resides. The host can be a
	// server or virtual machine.
	HostName *string

	// Time since agent or connector health was reported.
	LastHealthPingTime *string

	// Agent's first registration timestamp in UTC.
	RegisteredTime *string

	// The agent or connector version.
	Version *string
}

// Network details about the host where the agent/connector resides.
type AgentNetworkInfo struct {

	// The IP address for the host where the agent/connector resides.
	IpAddress *string

	// The MAC address for the host where the agent/connector resides.
	MacAddress *string
}

// Error messages returned for each import task that you deleted as a response for
// this command.
type BatchDeleteImportDataError struct {

	// The type of error that occurred for a specific import task.
	ErrorCode BatchDeleteImportDataErrorCode

	// The description of the error that occurred for a specific import task.
	ErrorDescription *string

	// The unique import ID associated with the error that occurred.
	ImportTaskId *string
}

// Tags for a configuration item. Tags are metadata that help you categorize IT
// assets.
type ConfigurationTag struct {

	// The configuration ID for the item to tag. You can specify a list of keys and
	// values.
	ConfigurationId *string

	// A type of IT asset to tag.
	ConfigurationType ConfigurationItemType

	// A type of tag on which to filter. For example, serverType.
	Key *string

	// The time the configuration tag was created in Coordinated Universal Time (UTC).
	TimeOfCreation *time.Time

	// A value on which to filter. For example key = serverType and value = web server.
	Value *string
}

// A list of continuous export descriptions.
type ContinuousExportDescription struct {

	// The type of data collector used to gather this data (currently only offered for
	// AGENT).
	DataSource DataSource

	// The unique ID assigned to this export.
	ExportId *string

	// The name of the s3 bucket where the export data parquet files are stored.
	S3Bucket *string

	// An object which describes how the data is stored.
	//
	//     * databaseName - the name
	// of the Glue database used to store the schema.
	SchemaStorageConfig map[string]*string

	// The timestamp representing when the continuous export was started.
	StartTime *time.Time

	// Describes the status of the export. Can be one of the following values:
	//
	//     *
	// START_IN_PROGRESS - setting up resources to start continuous export.
	//
	//     *
	// START_FAILED - an error occurred setting up continuous export. To recover, call
	// start-continuous-export again.
	//
	//     * ACTIVE - data is being exported to the
	// customer bucket.
	//
	//     * ERROR - an error occurred during export. To fix the
	// issue, call stop-continuous-export and start-continuous-export.
	//
	//     *
	// STOP_IN_PROGRESS - stopping the export.
	//
	//     * STOP_FAILED - an error occurred
	// stopping the export. To recover, call stop-continuous-export again.
	//
	//     *
	// INACTIVE - the continuous export has been stopped. Data is no longer being
	// exported to the customer bucket.
	Status ContinuousExportStatus

	// Contains information about any errors that have occurred. This data type can
	// have the following values:  <ul> <li> <p>ACCESS_DENIED - You don’t have
	// permission to start Data Exploration in Amazon Athena. Contact your AWS
	// administrator for help. For more information, see <a
	// href="http://docs.aws.amazon.com/application-discovery/latest/userguide/setting-up.html">Setting
	// Up AWS Application Discovery Service</a> in the Application Discovery Service
	// User Guide.</p> </li> <li> <p>DELIVERY_STREAM_LIMIT_FAILURE - You reached the
	// limit for Amazon Kinesis Data Firehose delivery streams. Reduce the number of
	// streams or request a limit increase and try again. For more information, see <a
	// href="http://docs.aws.amazon.com/streams/latest/dev/service-sizes-and-limits.html">Kinesis
	// Data Streams Limits</a> in the Amazon Kinesis Data Streams Developer Guide.</p>
	// </li> <li> <p>FIREHOSE_ROLE_MISSING - The Data Exploration feature is in an
	// error state because your IAM User is missing the
	// AWSApplicationDiscoveryServiceFirehose role. Turn on Data Exploration in Amazon
	// Athena and try again. For more information, see <a
	// href="http://docs.aws.amazon.com/application-discovery/latest/userguide/setting-up.html#setting-up-user-policy">Step
	// 3: Provide Application Discovery Service Access to Non-Administrator Users by
	// Attaching Policies</a> in the Application Discovery Service User Guide.</p>
	// </li> <li> <p>FIREHOSE_STREAM_DOES_NOT_EXIST - The Data Exploration feature is
	// in an error state because your IAM User is missing one or more of the Kinesis
	// data delivery streams.</p> </li> <li> <p>INTERNAL_FAILURE - The Data Exploration
	// feature is in an error state because of an internal failure. Try again later. If
	// this problem persists, contact AWS Support.</p> </li> <li>
	// <p>S3_BUCKET_LIMIT_FAILURE - You reached the limit for Amazon S3 buckets. Reduce
	// the number of Amazon S3 buckets or request a limit increase and try again. For
	// more information, see <a
	// href="http://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html">Bucket
	// Restrictions and Limitations</a> in the Amazon Simple Storage Service Developer
	// Guide.</p> </li> <li> <p>S3_NOT_SIGNED_UP - Your account is not signed up for
	// the Amazon S3 service. You must sign up before you can use Amazon S3. You can
	// sign up at the following URL: <a
	// href="https://aws.amazon.com/s3">https://aws.amazon.com/s3</a>.</p> </li> </ul>
	StatusDetail *string

	// The timestamp that represents when this continuous export was stopped.
	StopTime *time.Time
}

// Inventory data for installed discovery agents.
type CustomerAgentInfo struct {

	// Number of active discovery agents.
	//
	// This member is required.
	ActiveAgents *int32

	// Number of blacklisted discovery agents.
	//
	// This member is required.
	BlackListedAgents *int32

	// Number of healthy discovery agents
	//
	// This member is required.
	HealthyAgents *int32

	// Number of discovery agents with status SHUTDOWN.
	//
	// This member is required.
	ShutdownAgents *int32

	// Total number of discovery agents.
	//
	// This member is required.
	TotalAgents *int32

	// Number of unhealthy discovery agents.
	//
	// This member is required.
	UnhealthyAgents *int32

	// Number of unknown discovery agents.
	//
	// This member is required.
	UnknownAgents *int32
}

// Inventory data for installed discovery connectors.
type CustomerConnectorInfo struct {

	// Number of active discovery connectors.
	//
	// This member is required.
	ActiveConnectors *int32

	// Number of blacklisted discovery connectors.
	//
	// This member is required.
	BlackListedConnectors *int32

	// Number of healthy discovery connectors.
	//
	// This member is required.
	HealthyConnectors *int32

	// Number of discovery connectors with status SHUTDOWN,
	//
	// This member is required.
	ShutdownConnectors *int32

	// Total number of discovery connectors.
	//
	// This member is required.
	TotalConnectors *int32

	// Number of unhealthy discovery connectors.
	//
	// This member is required.
	UnhealthyConnectors *int32

	// Number of unknown discovery connectors.
	//
	// This member is required.
	UnknownConnectors *int32
}

// Used to select which agent's data is to be exported. A single agent ID may be
// selected for export using the StartExportTask
// (http://docs.aws.amazon.com/application-discovery/latest/APIReference/API_StartExportTask.html)
// action.
type ExportFilter struct {

	// Supported condition: EQUALS
	//
	// This member is required.
	Condition *string

	// A single ExportFilter name. Supported filters: agentId.
	//
	// This member is required.
	Name *string

	// A single agentId for a Discovery Agent. An agentId can be found using the
	// DescribeAgents
	// (http://docs.aws.amazon.com/application-discovery/latest/APIReference/API_DescribeExportTasks.html)
	// action. Typically an ADS agentId is in the form o-0123456789abcdef0.
	//
	// This member is required.
	Values []*string
}

// Information regarding the export status of discovered data. The value is an
// array of objects.
type ExportInfo struct {

	// A unique identifier used to query an export.
	//
	// This member is required.
	ExportId *string

	// The time that the data export was initiated.
	//
	// This member is required.
	ExportRequestTime *time.Time

	// The status of the data export job.
	//
	// This member is required.
	ExportStatus ExportStatus

	// A status message provided for API callers.
	//
	// This member is required.
	StatusMessage *string

	// A URL for an Amazon S3 bucket where you can review the exported data. The URL is
	// displayed only if the export succeeded.
	ConfigurationsDownloadUrl *string

	// If true, the export of agent information exceeded the size limit for a single
	// export and the exported data is incomplete for the requested time range. To
	// address this, select a smaller time range for the export by using startDate and
	// endDate.
	IsTruncated *bool

	// The endTime used in the StartExportTask request. If no endTime was requested,
	// this result does not appear in ExportInfo.
	RequestedEndTime *time.Time

	// The value of startTime parameter in the StartExportTask request. If no startTime
	// was requested, this result does not appear in ExportInfo.
	RequestedStartTime *time.Time
}

// A filter that can use conditional operators. For more information about filters,
// see Querying Discovered Configuration Items
// (https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-api-queries.html)
// in the AWS Application Discovery Service User Guide.
type Filter struct {

	// A conditional operator. The following operators are valid: EQUALS, NOT_EQUALS,
	// CONTAINS, NOT_CONTAINS. If you specify multiple filters, the system utilizes all
	// filters as though concatenated by AND. If you specify multiple values for a
	// particular filter, the system differentiates the values using OR. Calling either
	// DescribeConfigurations or ListConfigurations returns attributes of matching
	// configuration items.
	//
	// This member is required.
	Condition *string

	// The name of the filter.
	//
	// This member is required.
	Name *string

	// A string value on which to filter. For example, if you choose the
	// destinationServer.osVersion filter name, you could specify Ubuntu for the value.
	//
	// This member is required.
	Values []*string
}

// An array of information related to the import task request that includes status
// information, times, IDs, the Amazon S3 Object URL for the import file, and more.
type ImportTask struct {

	// The total number of application records in the import file that failed to be
	// imported.
	ApplicationImportFailure *int32

	// The total number of application records in the import file that were
	// successfully imported.
	ApplicationImportSuccess *int32

	// A unique token used to prevent the same import request from occurring more than
	// once. If you didn't provide a token, a token was automatically generated when
	// the import task request was sent.
	ClientRequestToken *string

	// A link to a compressed archive folder (in the ZIP format) that contains an error
	// log and a file of failed records. You can use these two files to quickly
	// identify records that failed, why they failed, and correct those records.
	// Afterward, you can upload the corrected file to your Amazon S3 bucket and create
	// another import task request.  <p>This field also includes authorization
	// information so you can confirm the authenticity of the compressed archive before
	// you download it.</p> <p>If some records failed to be imported we recommend that
	// you correct the records in the failed entries file and then imports that failed
	// entries file. This prevents you from having to correct and update the larger
	// original file and attempt importing it again.</p>
	ErrorsAndFailedEntriesZip *string

	// The time that the import task request finished, presented in the Unix time stamp
	// format.
	ImportCompletionTime *time.Time

	// The time that the import task request was deleted, presented in the Unix time
	// stamp format.
	ImportDeletedTime *time.Time

	// The time that the import task request was made, presented in the Unix time stamp
	// format.
	ImportRequestTime *time.Time

	// The unique ID for a specific import task. These IDs aren't globally unique, but
	// they are unique within an AWS account.
	ImportTaskId *string

	// The URL for your import file that you've uploaded to Amazon S3.
	ImportUrl *string

	// A descriptive name for an import task. You can use this name to filter future
	// requests related to this import task, such as identifying applications and
	// servers that were included in this import task. We recommend that you use a
	// meaningful name for each import task.
	Name *string

	// The total number of server records in the import file that failed to be
	// imported.
	ServerImportFailure *int32

	// The total number of server records in the import file that were successfully
	// imported.
	ServerImportSuccess *int32

	// The status of the import task. An import can have the status of IMPORT_COMPLETE
	// and still have some records fail to import from the overall request. More
	// information can be found in the downloadable archive defined in the
	// errorsAndFailedEntriesZip field, or in the Migration Hub management console.
	Status ImportStatus
}

// A name-values pair of elements you can use to filter the results when querying
// your import tasks. Currently, wildcards are not supported for filters.  <note>
// <p>When filtering by import status, all other filter values are ignored.</p>
// </note>
type ImportTaskFilter struct {

	// The name, status, or import task ID for a specific import task.
	Name ImportTaskFilterName

	// An array of strings that you can provide to match against a specific name,
	// status, or import task ID to filter the results for your import task queries.
	Values []*string
}

// Details about neighboring servers.
type NeighborConnectionDetail struct {

	// The number of open network connections with the neighboring server.
	//
	// This member is required.
	ConnectionsCount *int64

	// The ID of the server that accepted the network connection.
	//
	// This member is required.
	DestinationServerId *string

	// The ID of the server that opened the network connection.
	//
	// This member is required.
	SourceServerId *string

	// The destination network port for the connection.
	DestinationPort *int32

	// The network protocol used for the connection.
	TransportProtocol *string
}

// A field and direction for ordered output.
type OrderByElement struct {

	// The field on which to order.
	//
	// This member is required.
	FieldName *string

	// Ordering direction.
	SortOrder OrderString
}

// Metadata that help you categorize IT assets.
type Tag struct {

	// The type of tag on which to filter.
	//
	// This member is required.
	Key *string

	// A value for a tag key on which to filter.
	//
	// This member is required.
	Value *string
}

// The tag filter. Valid names are: tagKey, tagValue, configurationId.
type TagFilter struct {

	// A name of the tag filter.
	//
	// This member is required.
	Name *string

	// Values for the tag filter.
	//
	// This member is required.
	Values []*string
}