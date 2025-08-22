package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SubtitlesOutput represents the SubtitlesOutput schema from the OpenAPI specification
type SubtitlesOutput struct {
	Formats interface{} `json:"Formats,omitempty"`
	Outputstartindex interface{} `json:"OutputStartIndex,omitempty"`
	Subtitlefileuris interface{} `json:"SubtitleFileUris,omitempty"`
}

// ToxicityDetectionSettings represents the ToxicityDetectionSettings schema from the OpenAPI specification
type ToxicityDetectionSettings struct {
	Toxicitycategories interface{} `json:"ToxicityCategories"`
}

// StartTranscriptionJobResponse represents the StartTranscriptionJobResponse schema from the OpenAPI specification
type StartTranscriptionJobResponse struct {
	Transcriptionjob interface{} `json:"TranscriptionJob,omitempty"`
}

// GetVocabularyResponse represents the GetVocabularyResponse schema from the OpenAPI specification
type GetVocabularyResponse struct {
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Vocabularyname interface{} `json:"VocabularyName,omitempty"`
	Vocabularystate interface{} `json:"VocabularyState,omitempty"`
	Downloaduri interface{} `json:"DownloadUri,omitempty"`
}

// CreateCallAnalyticsCategoryResponse represents the CreateCallAnalyticsCategoryResponse schema from the OpenAPI specification
type CreateCallAnalyticsCategoryResponse struct {
	Categoryproperties interface{} `json:"CategoryProperties,omitempty"`
}

// GetVocabularyFilterResponse represents the GetVocabularyFilterResponse schema from the OpenAPI specification
type GetVocabularyFilterResponse struct {
	Vocabularyfiltername interface{} `json:"VocabularyFilterName,omitempty"`
	Downloaduri interface{} `json:"DownloadUri,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// CreateMedicalVocabularyRequest represents the CreateMedicalVocabularyRequest schema from the OpenAPI specification
type CreateMedicalVocabularyRequest struct {
	Languagecode interface{} `json:"LanguageCode"`
	Tags interface{} `json:"Tags,omitempty"`
	Vocabularyfileuri interface{} `json:"VocabularyFileUri"`
	Vocabularyname interface{} `json:"VocabularyName"`
}

// DeleteLanguageModelRequest represents the DeleteLanguageModelRequest schema from the OpenAPI specification
type DeleteLanguageModelRequest struct {
	Modelname interface{} `json:"ModelName"`
}

// DeleteTranscriptionJobRequest represents the DeleteTranscriptionJobRequest schema from the OpenAPI specification
type DeleteTranscriptionJobRequest struct {
	Transcriptionjobname interface{} `json:"TranscriptionJobName"`
}

// GetCallAnalyticsJobResponse represents the GetCallAnalyticsJobResponse schema from the OpenAPI specification
type GetCallAnalyticsJobResponse struct {
	Callanalyticsjob interface{} `json:"CallAnalyticsJob,omitempty"`
}

// ListMedicalVocabulariesRequest represents the ListMedicalVocabulariesRequest schema from the OpenAPI specification
type ListMedicalVocabulariesRequest struct {
	Stateequals interface{} `json:"StateEquals,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// LanguageIdSettings represents the LanguageIdSettings schema from the OpenAPI specification
type LanguageIdSettings struct {
	Vocabularyfiltername interface{} `json:"VocabularyFilterName,omitempty"`
	Vocabularyname interface{} `json:"VocabularyName,omitempty"`
	Languagemodelname interface{} `json:"LanguageModelName,omitempty"`
}

// ListCallAnalyticsJobsResponse represents the ListCallAnalyticsJobsResponse schema from the OpenAPI specification
type ListCallAnalyticsJobsResponse struct {
	Status interface{} `json:"Status,omitempty"`
	Callanalyticsjobsummaries interface{} `json:"CallAnalyticsJobSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
	Tagkeys interface{} `json:"TagKeys"`
}

// InputDataConfig represents the InputDataConfig schema from the OpenAPI specification
type InputDataConfig struct {
	S3uri interface{} `json:"S3Uri"`
	Tuningdatas3uri interface{} `json:"TuningDataS3Uri,omitempty"`
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn"`
}

// GetCallAnalyticsCategoryResponse represents the GetCallAnalyticsCategoryResponse schema from the OpenAPI specification
type GetCallAnalyticsCategoryResponse struct {
	Categoryproperties interface{} `json:"CategoryProperties,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Value interface{} `json:"Value"`
	Key interface{} `json:"Key"`
}

// MedicalTranscriptionSetting represents the MedicalTranscriptionSetting schema from the OpenAPI specification
type MedicalTranscriptionSetting struct {
	Vocabularyname interface{} `json:"VocabularyName,omitempty"`
	Channelidentification interface{} `json:"ChannelIdentification,omitempty"`
	Maxalternatives interface{} `json:"MaxAlternatives,omitempty"`
	Maxspeakerlabels interface{} `json:"MaxSpeakerLabels,omitempty"`
	Showalternatives interface{} `json:"ShowAlternatives,omitempty"`
	Showspeakerlabels interface{} `json:"ShowSpeakerLabels,omitempty"`
}

// ListVocabulariesResponse represents the ListVocabulariesResponse schema from the OpenAPI specification
type ListVocabulariesResponse struct {
	Status interface{} `json:"Status,omitempty"`
	Vocabularies interface{} `json:"Vocabularies,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// KMSEncryptionContextMap represents the KMSEncryptionContextMap schema from the OpenAPI specification
type KMSEncryptionContextMap struct {
}

// GetCallAnalyticsCategoryRequest represents the GetCallAnalyticsCategoryRequest schema from the OpenAPI specification
type GetCallAnalyticsCategoryRequest struct {
	Categoryname interface{} `json:"CategoryName"`
}

// StartMedicalTranscriptionJobResponse represents the StartMedicalTranscriptionJobResponse schema from the OpenAPI specification
type StartMedicalTranscriptionJobResponse struct {
	Medicaltranscriptionjob interface{} `json:"MedicalTranscriptionJob,omitempty"`
}

// DeleteCallAnalyticsJobRequest represents the DeleteCallAnalyticsJobRequest schema from the OpenAPI specification
type DeleteCallAnalyticsJobRequest struct {
	Callanalyticsjobname interface{} `json:"CallAnalyticsJobName"`
}

// TranscriptionJob represents the TranscriptionJob schema from the OpenAPI specification
type TranscriptionJob struct {
	Subtitles interface{} `json:"Subtitles,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Contentredaction interface{} `json:"ContentRedaction,omitempty"`
	Mediasampleratehertz interface{} `json:"MediaSampleRateHertz,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Transcriptionjobstatus interface{} `json:"TranscriptionJobStatus,omitempty"`
	Jobexecutionsettings interface{} `json:"JobExecutionSettings,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Mediaformat interface{} `json:"MediaFormat,omitempty"`
	Languageoptions interface{} `json:"LanguageOptions,omitempty"`
	Languagecodes interface{} `json:"LanguageCodes,omitempty"`
	Toxicitydetection interface{} `json:"ToxicityDetection,omitempty"`
	Media interface{} `json:"Media,omitempty"`
	Settings interface{} `json:"Settings,omitempty"`
	Identifylanguage interface{} `json:"IdentifyLanguage,omitempty"`
	Languageidsettings interface{} `json:"LanguageIdSettings,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Transcriptionjobname interface{} `json:"TranscriptionJobName,omitempty"`
	Completiontime interface{} `json:"CompletionTime,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Identifiedlanguagescore interface{} `json:"IdentifiedLanguageScore,omitempty"`
	Identifymultiplelanguages interface{} `json:"IdentifyMultipleLanguages,omitempty"`
	Modelsettings interface{} `json:"ModelSettings,omitempty"`
	Transcript interface{} `json:"Transcript,omitempty"`
}

// GetMedicalTranscriptionJobRequest represents the GetMedicalTranscriptionJobRequest schema from the OpenAPI specification
type GetMedicalTranscriptionJobRequest struct {
	Medicaltranscriptionjobname interface{} `json:"MedicalTranscriptionJobName"`
}

// UpdateCallAnalyticsCategoryResponse represents the UpdateCallAnalyticsCategoryResponse schema from the OpenAPI specification
type UpdateCallAnalyticsCategoryResponse struct {
	Categoryproperties interface{} `json:"CategoryProperties,omitempty"`
}

// GetVocabularyRequest represents the GetVocabularyRequest schema from the OpenAPI specification
type GetVocabularyRequest struct {
	Vocabularyname interface{} `json:"VocabularyName"`
}

// ListVocabularyFiltersRequest represents the ListVocabularyFiltersRequest schema from the OpenAPI specification
type ListVocabularyFiltersRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CallAnalyticsJob represents the CallAnalyticsJob schema from the OpenAPI specification
type CallAnalyticsJob struct {
	Callanalyticsjobname interface{} `json:"CallAnalyticsJobName,omitempty"`
	Channeldefinitions interface{} `json:"ChannelDefinitions,omitempty"`
	Identifiedlanguagescore interface{} `json:"IdentifiedLanguageScore,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Transcript Transcript `json:"Transcript,omitempty"` // Provides you with the Amazon S3 URI you can use to access your transcript.
	Callanalyticsjobstatus interface{} `json:"CallAnalyticsJobStatus,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Mediaformat interface{} `json:"MediaFormat,omitempty"`
	Settings interface{} `json:"Settings,omitempty"`
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Completiontime interface{} `json:"CompletionTime,omitempty"`
	Media interface{} `json:"Media,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Mediasampleratehertz interface{} `json:"MediaSampleRateHertz,omitempty"`
}

// CreateMedicalVocabularyResponse represents the CreateMedicalVocabularyResponse schema from the OpenAPI specification
type CreateMedicalVocabularyResponse struct {
	Vocabularystate interface{} `json:"VocabularyState,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Vocabularyname interface{} `json:"VocabularyName,omitempty"`
}

// SentimentFilter represents the SentimentFilter schema from the OpenAPI specification
type SentimentFilter struct {
	Absolutetimerange interface{} `json:"AbsoluteTimeRange,omitempty"`
	Negate interface{} `json:"Negate,omitempty"`
	Participantrole interface{} `json:"ParticipantRole,omitempty"`
	Relativetimerange interface{} `json:"RelativeTimeRange,omitempty"`
	Sentiments interface{} `json:"Sentiments"`
}

// NonTalkTimeFilter represents the NonTalkTimeFilter schema from the OpenAPI specification
type NonTalkTimeFilter struct {
	Relativetimerange interface{} `json:"RelativeTimeRange,omitempty"`
	Threshold interface{} `json:"Threshold,omitempty"`
	Absolutetimerange interface{} `json:"AbsoluteTimeRange,omitempty"`
	Negate interface{} `json:"Negate,omitempty"`
}

// Subtitles represents the Subtitles schema from the OpenAPI specification
type Subtitles struct {
	Formats interface{} `json:"Formats,omitempty"`
	Outputstartindex interface{} `json:"OutputStartIndex,omitempty"`
}

// DeleteCallAnalyticsJobResponse represents the DeleteCallAnalyticsJobResponse schema from the OpenAPI specification
type DeleteCallAnalyticsJobResponse struct {
}

// UpdateVocabularyFilterResponse represents the UpdateVocabularyFilterResponse schema from the OpenAPI specification
type UpdateVocabularyFilterResponse struct {
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Vocabularyfiltername interface{} `json:"VocabularyFilterName,omitempty"`
}

// MedicalTranscript represents the MedicalTranscript schema from the OpenAPI specification
type MedicalTranscript struct {
	Transcriptfileuri interface{} `json:"TranscriptFileUri,omitempty"`
}

// ListCallAnalyticsJobsRequest represents the ListCallAnalyticsJobsRequest schema from the OpenAPI specification
type ListCallAnalyticsJobsRequest struct {
	Status interface{} `json:"Status,omitempty"`
	Jobnamecontains interface{} `json:"JobNameContains,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// StartTranscriptionJobRequest represents the StartTranscriptionJobRequest schema from the OpenAPI specification
type StartTranscriptionJobRequest struct {
	Outputencryptionkmskeyid interface{} `json:"OutputEncryptionKMSKeyId,omitempty"`
	Identifylanguage interface{} `json:"IdentifyLanguage,omitempty"`
	Mediasampleratehertz interface{} `json:"MediaSampleRateHertz,omitempty"`
	Kmsencryptioncontext interface{} `json:"KMSEncryptionContext,omitempty"`
	Outputbucketname interface{} `json:"OutputBucketName,omitempty"`
	Identifymultiplelanguages interface{} `json:"IdentifyMultipleLanguages,omitempty"`
	Jobexecutionsettings interface{} `json:"JobExecutionSettings,omitempty"`
	Mediaformat interface{} `json:"MediaFormat,omitempty"`
	Settings interface{} `json:"Settings,omitempty"`
	Outputkey interface{} `json:"OutputKey,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Subtitles interface{} `json:"Subtitles,omitempty"`
	Media interface{} `json:"Media"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Languageoptions interface{} `json:"LanguageOptions,omitempty"`
	Toxicitydetection interface{} `json:"ToxicityDetection,omitempty"`
	Contentredaction interface{} `json:"ContentRedaction,omitempty"`
	Languageidsettings interface{} `json:"LanguageIdSettings,omitempty"`
	Modelsettings interface{} `json:"ModelSettings,omitempty"`
	Transcriptionjobname interface{} `json:"TranscriptionJobName"`
}

// UpdateVocabularyResponse represents the UpdateVocabularyResponse schema from the OpenAPI specification
type UpdateVocabularyResponse struct {
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Vocabularyname interface{} `json:"VocabularyName,omitempty"`
	Vocabularystate interface{} `json:"VocabularyState,omitempty"`
}

// DeleteCallAnalyticsCategoryResponse represents the DeleteCallAnalyticsCategoryResponse schema from the OpenAPI specification
type DeleteCallAnalyticsCategoryResponse struct {
}

// UpdateCallAnalyticsCategoryRequest represents the UpdateCallAnalyticsCategoryRequest schema from the OpenAPI specification
type UpdateCallAnalyticsCategoryRequest struct {
	Categoryname interface{} `json:"CategoryName"`
	Inputtype interface{} `json:"InputType,omitempty"`
	Rules interface{} `json:"Rules"`
}

// CallAnalyticsJobSettings represents the CallAnalyticsJobSettings schema from the OpenAPI specification
type CallAnalyticsJobSettings struct {
	Languageoptions interface{} `json:"LanguageOptions,omitempty"`
	Vocabularyfiltermethod interface{} `json:"VocabularyFilterMethod,omitempty"`
	Vocabularyfiltername interface{} `json:"VocabularyFilterName,omitempty"`
	Vocabularyname interface{} `json:"VocabularyName,omitempty"`
	Contentredaction ContentRedaction `json:"ContentRedaction,omitempty"` // Makes it possible to redact or flag specified personally identifiable information (PII) in your transcript. If you use <code>ContentRedaction</code>, you must also include the sub-parameters: <code>PiiEntityTypes</code>, <code>RedactionOutput</code>, and <code>RedactionType</code>.
	Languageidsettings interface{} `json:"LanguageIdSettings,omitempty"`
	Languagemodelname interface{} `json:"LanguageModelName,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// LanguageModel represents the LanguageModel schema from the OpenAPI specification
type LanguageModel struct {
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Modelname interface{} `json:"ModelName,omitempty"`
	Createtime interface{} `json:"CreateTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Modelstatus interface{} `json:"ModelStatus,omitempty"`
	Upgradeavailability interface{} `json:"UpgradeAvailability,omitempty"`
	Basemodelname interface{} `json:"BaseModelName,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig,omitempty"`
}

// DeleteVocabularyFilterRequest represents the DeleteVocabularyFilterRequest schema from the OpenAPI specification
type DeleteVocabularyFilterRequest struct {
	Vocabularyfiltername interface{} `json:"VocabularyFilterName"`
}

// ListMedicalTranscriptionJobsRequest represents the ListMedicalTranscriptionJobsRequest schema from the OpenAPI specification
type ListMedicalTranscriptionJobsRequest struct {
	Status interface{} `json:"Status,omitempty"`
	Jobnamecontains interface{} `json:"JobNameContains,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ModelSettings represents the ModelSettings schema from the OpenAPI specification
type ModelSettings struct {
	Languagemodelname interface{} `json:"LanguageModelName,omitempty"`
}

// ContentRedaction represents the ContentRedaction schema from the OpenAPI specification
type ContentRedaction struct {
	Redactionoutput interface{} `json:"RedactionOutput"`
	Redactiontype interface{} `json:"RedactionType"`
	Piientitytypes interface{} `json:"PiiEntityTypes,omitempty"`
}

// CreateVocabularyFilterRequest represents the CreateVocabularyFilterRequest schema from the OpenAPI specification
type CreateVocabularyFilterRequest struct {
	Languagecode interface{} `json:"LanguageCode"`
	Tags interface{} `json:"Tags,omitempty"`
	Vocabularyfilterfileuri interface{} `json:"VocabularyFilterFileUri,omitempty"`
	Vocabularyfiltername interface{} `json:"VocabularyFilterName"`
	Words interface{} `json:"Words,omitempty"`
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn,omitempty"`
}

// UpdateVocabularyFilterRequest represents the UpdateVocabularyFilterRequest schema from the OpenAPI specification
type UpdateVocabularyFilterRequest struct {
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn,omitempty"`
	Vocabularyfilterfileuri interface{} `json:"VocabularyFilterFileUri,omitempty"`
	Vocabularyfiltername interface{} `json:"VocabularyFilterName"`
	Words interface{} `json:"Words,omitempty"`
}

// CreateCallAnalyticsCategoryRequest represents the CreateCallAnalyticsCategoryRequest schema from the OpenAPI specification
type CreateCallAnalyticsCategoryRequest struct {
	Categoryname interface{} `json:"CategoryName"`
	Inputtype interface{} `json:"InputType,omitempty"`
	Rules interface{} `json:"Rules"`
}

// LanguageIdSettingsMap represents the LanguageIdSettingsMap schema from the OpenAPI specification
type LanguageIdSettingsMap struct {
}

// Rule represents the Rule schema from the OpenAPI specification
type Rule struct {
	Sentimentfilter interface{} `json:"SentimentFilter,omitempty"`
	Transcriptfilter interface{} `json:"TranscriptFilter,omitempty"`
	Interruptionfilter interface{} `json:"InterruptionFilter,omitempty"`
	Nontalktimefilter interface{} `json:"NonTalkTimeFilter,omitempty"`
}

// Media represents the Media schema from the OpenAPI specification
type Media struct {
	Mediafileuri interface{} `json:"MediaFileUri,omitempty"`
	Redactedmediafileuri interface{} `json:"RedactedMediaFileUri,omitempty"`
}

// VocabularyInfo represents the VocabularyInfo schema from the OpenAPI specification
type VocabularyInfo struct {
	Vocabularystate interface{} `json:"VocabularyState,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Vocabularyname interface{} `json:"VocabularyName,omitempty"`
}

// TranscriptFilter represents the TranscriptFilter schema from the OpenAPI specification
type TranscriptFilter struct {
	Participantrole interface{} `json:"ParticipantRole,omitempty"`
	Relativetimerange interface{} `json:"RelativeTimeRange,omitempty"`
	Targets interface{} `json:"Targets"`
	Transcriptfiltertype interface{} `json:"TranscriptFilterType"`
	Absolutetimerange interface{} `json:"AbsoluteTimeRange,omitempty"`
	Negate interface{} `json:"Negate,omitempty"`
}

// StartCallAnalyticsJobResponse represents the StartCallAnalyticsJobResponse schema from the OpenAPI specification
type StartCallAnalyticsJobResponse struct {
	Callanalyticsjob interface{} `json:"CallAnalyticsJob,omitempty"`
}

// CreateVocabularyResponse represents the CreateVocabularyResponse schema from the OpenAPI specification
type CreateVocabularyResponse struct {
	Vocabularyname interface{} `json:"VocabularyName,omitempty"`
	Vocabularystate interface{} `json:"VocabularyState,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// MedicalTranscriptionJob represents the MedicalTranscriptionJob schema from the OpenAPI specification
type MedicalTranscriptionJob struct {
	Settings interface{} `json:"Settings,omitempty"`
	Transcriptionjobstatus interface{} `json:"TranscriptionJobStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Transcript interface{} `json:"Transcript,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Medicaltranscriptionjobname interface{} `json:"MedicalTranscriptionJobName,omitempty"`
	Contentidentificationtype interface{} `json:"ContentIdentificationType,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Completiontime interface{} `json:"CompletionTime,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Mediaformat interface{} `json:"MediaFormat,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Media Media `json:"Media,omitempty"` // <p>Describes the Amazon S3 location of the media file you want to use in your request.</p> <p>For information on supported media formats, refer to the <a href="https://docs.aws.amazon.com/APIReference/API_StartTranscriptionJob.html#transcribe-StartTranscriptionJob-request-MediaFormat">MediaFormat</a> parameter or the <a href="https://docs.aws.amazon.com/transcribe/latest/dg/how-input.html#how-input-audio">Media formats</a> section in the Amazon S3 Developer Guide.</p>
	Mediasampleratehertz interface{} `json:"MediaSampleRateHertz,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Specialty interface{} `json:"Specialty,omitempty"`
}

// ListTranscriptionJobsRequest represents the ListTranscriptionJobsRequest schema from the OpenAPI specification
type ListTranscriptionJobsRequest struct {
	Status interface{} `json:"Status,omitempty"`
	Jobnamecontains interface{} `json:"JobNameContains,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Transcript represents the Transcript schema from the OpenAPI specification
type Transcript struct {
	Redactedtranscriptfileuri interface{} `json:"RedactedTranscriptFileUri,omitempty"`
	Transcriptfileuri interface{} `json:"TranscriptFileUri,omitempty"`
}

// StartCallAnalyticsJobRequest represents the StartCallAnalyticsJobRequest schema from the OpenAPI specification
type StartCallAnalyticsJobRequest struct {
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn,omitempty"`
	Media interface{} `json:"Media"`
	Outputencryptionkmskeyid interface{} `json:"OutputEncryptionKMSKeyId,omitempty"`
	Outputlocation interface{} `json:"OutputLocation,omitempty"`
	Settings interface{} `json:"Settings,omitempty"`
	Callanalyticsjobname interface{} `json:"CallAnalyticsJobName"`
	Channeldefinitions interface{} `json:"ChannelDefinitions,omitempty"`
}

// TranscriptionJobSummary represents the TranscriptionJobSummary schema from the OpenAPI specification
type TranscriptionJobSummary struct {
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Languagecodes interface{} `json:"LanguageCodes,omitempty"`
	Modelsettings ModelSettings `json:"ModelSettings,omitempty"` // <p>Provides the name of the custom language model that was included in the specified transcription job.</p> <p>Only use <code>ModelSettings</code> with the <code>LanguageModelName</code> sub-parameter if you're <b>not</b> using automatic language identification (<code/>). If using <code>LanguageIdSettings</code> in your request, this parameter contains a <code>LanguageModelName</code> sub-parameter.</p>
	Completiontime interface{} `json:"CompletionTime,omitempty"`
	Toxicitydetection interface{} `json:"ToxicityDetection,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Identifiedlanguagescore interface{} `json:"IdentifiedLanguageScore,omitempty"`
	Identifymultiplelanguages interface{} `json:"IdentifyMultipleLanguages,omitempty"`
	Outputlocationtype interface{} `json:"OutputLocationType,omitempty"`
	Identifylanguage interface{} `json:"IdentifyLanguage,omitempty"`
	Contentredaction interface{} `json:"ContentRedaction,omitempty"`
	Transcriptionjobname interface{} `json:"TranscriptionJobName,omitempty"`
	Transcriptionjobstatus interface{} `json:"TranscriptionJobStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
}

// GetTranscriptionJobResponse represents the GetTranscriptionJobResponse schema from the OpenAPI specification
type GetTranscriptionJobResponse struct {
	Transcriptionjob interface{} `json:"TranscriptionJob,omitempty"`
}

// ListLanguageModelsResponse represents the ListLanguageModelsResponse schema from the OpenAPI specification
type ListLanguageModelsResponse struct {
	Models interface{} `json:"Models,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CallAnalyticsJobSummary represents the CallAnalyticsJobSummary schema from the OpenAPI specification
type CallAnalyticsJobSummary struct {
	Callanalyticsjobname interface{} `json:"CallAnalyticsJobName,omitempty"`
	Callanalyticsjobstatus interface{} `json:"CallAnalyticsJobStatus,omitempty"`
	Completiontime interface{} `json:"CompletionTime,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
}

// ListTranscriptionJobsResponse represents the ListTranscriptionJobsResponse schema from the OpenAPI specification
type ListTranscriptionJobsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Transcriptionjobsummaries interface{} `json:"TranscriptionJobSummaries,omitempty"`
}

// CreateLanguageModelRequest represents the CreateLanguageModelRequest schema from the OpenAPI specification
type CreateLanguageModelRequest struct {
	Basemodelname interface{} `json:"BaseModelName"`
	Inputdataconfig interface{} `json:"InputDataConfig"`
	Languagecode interface{} `json:"LanguageCode"`
	Modelname interface{} `json:"ModelName"`
	Tags interface{} `json:"Tags,omitempty"`
}

// StartMedicalTranscriptionJobRequest represents the StartMedicalTranscriptionJobRequest schema from the OpenAPI specification
type StartMedicalTranscriptionJobRequest struct {
	Contentidentificationtype interface{} `json:"ContentIdentificationType,omitempty"`
	Kmsencryptioncontext interface{} `json:"KMSEncryptionContext,omitempty"`
	Languagecode interface{} `json:"LanguageCode"`
	Mediaformat interface{} `json:"MediaFormat,omitempty"`
	Medicaltranscriptionjobname interface{} `json:"MedicalTranscriptionJobName"`
	Outputkey interface{} `json:"OutputKey,omitempty"`
	Media Media `json:"Media"` // <p>Describes the Amazon S3 location of the media file you want to use in your request.</p> <p>For information on supported media formats, refer to the <a href="https://docs.aws.amazon.com/APIReference/API_StartTranscriptionJob.html#transcribe-StartTranscriptionJob-request-MediaFormat">MediaFormat</a> parameter or the <a href="https://docs.aws.amazon.com/transcribe/latest/dg/how-input.html#how-input-audio">Media formats</a> section in the Amazon S3 Developer Guide.</p>
	Outputencryptionkmskeyid interface{} `json:"OutputEncryptionKMSKeyId,omitempty"`
	Settings interface{} `json:"Settings,omitempty"`
	Specialty interface{} `json:"Specialty"`
	TypeField interface{} `json:"Type"`
	Mediasampleratehertz interface{} `json:"MediaSampleRateHertz,omitempty"`
	Outputbucketname interface{} `json:"OutputBucketName"`
	Tags interface{} `json:"Tags,omitempty"`
}

// GetMedicalVocabularyRequest represents the GetMedicalVocabularyRequest schema from the OpenAPI specification
type GetMedicalVocabularyRequest struct {
	Vocabularyname interface{} `json:"VocabularyName"`
}

// InterruptionFilter represents the InterruptionFilter schema from the OpenAPI specification
type InterruptionFilter struct {
	Participantrole interface{} `json:"ParticipantRole,omitempty"`
	Relativetimerange interface{} `json:"RelativeTimeRange,omitempty"`
	Threshold interface{} `json:"Threshold,omitempty"`
	Absolutetimerange interface{} `json:"AbsoluteTimeRange,omitempty"`
	Negate interface{} `json:"Negate,omitempty"`
}

// JobExecutionSettings represents the JobExecutionSettings schema from the OpenAPI specification
type JobExecutionSettings struct {
	Allowdeferredexecution interface{} `json:"AllowDeferredExecution,omitempty"`
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn,omitempty"`
}

// AbsoluteTimeRange represents the AbsoluteTimeRange schema from the OpenAPI specification
type AbsoluteTimeRange struct {
	Endtime interface{} `json:"EndTime,omitempty"`
	First interface{} `json:"First,omitempty"`
	Last interface{} `json:"Last,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// GetMedicalVocabularyResponse represents the GetMedicalVocabularyResponse schema from the OpenAPI specification
type GetMedicalVocabularyResponse struct {
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Vocabularyname interface{} `json:"VocabularyName,omitempty"`
	Vocabularystate interface{} `json:"VocabularyState,omitempty"`
	Downloaduri interface{} `json:"DownloadUri,omitempty"`
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
}

// CreateVocabularyFilterResponse represents the CreateVocabularyFilterResponse schema from the OpenAPI specification
type CreateVocabularyFilterResponse struct {
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Vocabularyfiltername interface{} `json:"VocabularyFilterName,omitempty"`
}

// GetMedicalTranscriptionJobResponse represents the GetMedicalTranscriptionJobResponse schema from the OpenAPI specification
type GetMedicalTranscriptionJobResponse struct {
	Medicaltranscriptionjob interface{} `json:"MedicalTranscriptionJob,omitempty"`
}

// ListVocabularyFiltersResponse represents the ListVocabularyFiltersResponse schema from the OpenAPI specification
type ListVocabularyFiltersResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Vocabularyfilters interface{} `json:"VocabularyFilters,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
}

// ListCallAnalyticsCategoriesResponse represents the ListCallAnalyticsCategoriesResponse schema from the OpenAPI specification
type ListCallAnalyticsCategoriesResponse struct {
	Categories interface{} `json:"Categories,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// LanguageCodeItem represents the LanguageCodeItem schema from the OpenAPI specification
type LanguageCodeItem struct {
	Durationinseconds interface{} `json:"DurationInSeconds,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
}

// ListLanguageModelsRequest represents the ListLanguageModelsRequest schema from the OpenAPI specification
type ListLanguageModelsRequest struct {
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Statusequals interface{} `json:"StatusEquals,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// RelativeTimeRange represents the RelativeTimeRange schema from the OpenAPI specification
type RelativeTimeRange struct {
	Startpercentage interface{} `json:"StartPercentage,omitempty"`
	Endpercentage interface{} `json:"EndPercentage,omitempty"`
	First interface{} `json:"First,omitempty"`
	Last interface{} `json:"Last,omitempty"`
}

// UpdateVocabularyRequest represents the UpdateVocabularyRequest schema from the OpenAPI specification
type UpdateVocabularyRequest struct {
	Phrases interface{} `json:"Phrases,omitempty"`
	Vocabularyfileuri interface{} `json:"VocabularyFileUri,omitempty"`
	Vocabularyname interface{} `json:"VocabularyName"`
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn,omitempty"`
	Languagecode interface{} `json:"LanguageCode"`
}

// DescribeLanguageModelRequest represents the DescribeLanguageModelRequest schema from the OpenAPI specification
type DescribeLanguageModelRequest struct {
	Modelname interface{} `json:"ModelName"`
}

// Settings represents the Settings schema from the OpenAPI specification
type Settings struct {
	Maxspeakerlabels interface{} `json:"MaxSpeakerLabels,omitempty"`
	Showalternatives interface{} `json:"ShowAlternatives,omitempty"`
	Showspeakerlabels interface{} `json:"ShowSpeakerLabels,omitempty"`
	Vocabularyfiltermethod interface{} `json:"VocabularyFilterMethod,omitempty"`
	Vocabularyfiltername interface{} `json:"VocabularyFilterName,omitempty"`
	Vocabularyname interface{} `json:"VocabularyName,omitempty"`
	Channelidentification interface{} `json:"ChannelIdentification,omitempty"`
	Maxalternatives interface{} `json:"MaxAlternatives,omitempty"`
}

// ListMedicalVocabulariesResponse represents the ListMedicalVocabulariesResponse schema from the OpenAPI specification
type ListMedicalVocabulariesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Vocabularies interface{} `json:"Vocabularies,omitempty"`
}

// CategoryProperties represents the CategoryProperties schema from the OpenAPI specification
type CategoryProperties struct {
	Createtime interface{} `json:"CreateTime,omitempty"`
	Inputtype interface{} `json:"InputType,omitempty"`
	Lastupdatetime interface{} `json:"LastUpdateTime,omitempty"`
	Rules interface{} `json:"Rules,omitempty"`
	Categoryname interface{} `json:"CategoryName,omitempty"`
}

// MedicalTranscriptionJobSummary represents the MedicalTranscriptionJobSummary schema from the OpenAPI specification
type MedicalTranscriptionJobSummary struct {
	Failurereason interface{} `json:"FailureReason,omitempty"`
	Outputlocationtype interface{} `json:"OutputLocationType,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Contentidentificationtype interface{} `json:"ContentIdentificationType,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Medicaltranscriptionjobname interface{} `json:"MedicalTranscriptionJobName,omitempty"`
	Specialty interface{} `json:"Specialty,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Completiontime interface{} `json:"CompletionTime,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Transcriptionjobstatus interface{} `json:"TranscriptionJobStatus,omitempty"`
}

// VocabularyFilterInfo represents the VocabularyFilterInfo schema from the OpenAPI specification
type VocabularyFilterInfo struct {
	Vocabularyfiltername interface{} `json:"VocabularyFilterName,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
}

// GetVocabularyFilterRequest represents the GetVocabularyFilterRequest schema from the OpenAPI specification
type GetVocabularyFilterRequest struct {
	Vocabularyfiltername interface{} `json:"VocabularyFilterName"`
}

// DeleteCallAnalyticsCategoryRequest represents the DeleteCallAnalyticsCategoryRequest schema from the OpenAPI specification
type DeleteCallAnalyticsCategoryRequest struct {
	Categoryname interface{} `json:"CategoryName"`
}

// DeleteVocabularyRequest represents the DeleteVocabularyRequest schema from the OpenAPI specification
type DeleteVocabularyRequest struct {
	Vocabularyname interface{} `json:"VocabularyName"`
}

// CreateVocabularyRequest represents the CreateVocabularyRequest schema from the OpenAPI specification
type CreateVocabularyRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Vocabularyfileuri interface{} `json:"VocabularyFileUri,omitempty"`
	Vocabularyname interface{} `json:"VocabularyName"`
	Dataaccessrolearn interface{} `json:"DataAccessRoleArn,omitempty"`
	Languagecode interface{} `json:"LanguageCode"`
	Phrases interface{} `json:"Phrases,omitempty"`
}

// UpdateMedicalVocabularyRequest represents the UpdateMedicalVocabularyRequest schema from the OpenAPI specification
type UpdateMedicalVocabularyRequest struct {
	Vocabularyfileuri interface{} `json:"VocabularyFileUri"`
	Vocabularyname interface{} `json:"VocabularyName"`
	Languagecode interface{} `json:"LanguageCode"`
}

// DeleteMedicalVocabularyRequest represents the DeleteMedicalVocabularyRequest schema from the OpenAPI specification
type DeleteMedicalVocabularyRequest struct {
	Vocabularyname interface{} `json:"VocabularyName"`
}

// ListMedicalTranscriptionJobsResponse represents the ListMedicalTranscriptionJobsResponse schema from the OpenAPI specification
type ListMedicalTranscriptionJobsResponse struct {
	Medicaltranscriptionjobsummaries interface{} `json:"MedicalTranscriptionJobSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// GetTranscriptionJobRequest represents the GetTranscriptionJobRequest schema from the OpenAPI specification
type GetTranscriptionJobRequest struct {
	Transcriptionjobname interface{} `json:"TranscriptionJobName"`
}

// UpdateMedicalVocabularyResponse represents the UpdateMedicalVocabularyResponse schema from the OpenAPI specification
type UpdateMedicalVocabularyResponse struct {
	Vocabularystate interface{} `json:"VocabularyState,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Vocabularyname interface{} `json:"VocabularyName,omitempty"`
}

// DeleteMedicalTranscriptionJobRequest represents the DeleteMedicalTranscriptionJobRequest schema from the OpenAPI specification
type DeleteMedicalTranscriptionJobRequest struct {
	Medicaltranscriptionjobname interface{} `json:"MedicalTranscriptionJobName"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Resourcearn interface{} `json:"ResourceArn,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ChannelDefinition represents the ChannelDefinition schema from the OpenAPI specification
type ChannelDefinition struct {
	Channelid interface{} `json:"ChannelId,omitempty"`
	Participantrole interface{} `json:"ParticipantRole,omitempty"`
}

// CreateLanguageModelResponse represents the CreateLanguageModelResponse schema from the OpenAPI specification
type CreateLanguageModelResponse struct {
	Modelname interface{} `json:"ModelName,omitempty"`
	Modelstatus interface{} `json:"ModelStatus,omitempty"`
	Basemodelname interface{} `json:"BaseModelName,omitempty"`
	Inputdataconfig interface{} `json:"InputDataConfig,omitempty"`
	Languagecode interface{} `json:"LanguageCode,omitempty"`
}

// ListVocabulariesRequest represents the ListVocabulariesRequest schema from the OpenAPI specification
type ListVocabulariesRequest struct {
	Stateequals interface{} `json:"StateEquals,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Namecontains interface{} `json:"NameContains,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetCallAnalyticsJobRequest represents the GetCallAnalyticsJobRequest schema from the OpenAPI specification
type GetCallAnalyticsJobRequest struct {
	Callanalyticsjobname interface{} `json:"CallAnalyticsJobName"`
}

// DescribeLanguageModelResponse represents the DescribeLanguageModelResponse schema from the OpenAPI specification
type DescribeLanguageModelResponse struct {
	Languagemodel interface{} `json:"LanguageModel,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// ListCallAnalyticsCategoriesRequest represents the ListCallAnalyticsCategoriesRequest schema from the OpenAPI specification
type ListCallAnalyticsCategoriesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}
