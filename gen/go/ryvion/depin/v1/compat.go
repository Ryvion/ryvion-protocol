package depinv1

const (
	DeviceClassUnspecified DeviceClass = DeviceClass_DEVICE_CLASS_UNSPECIFIED
	DeviceClassGpuCluster  DeviceClass = DeviceClass_DEVICE_CLASS_GPU_CLUSTER
	DeviceClassEdgeAI      DeviceClass = DeviceClass_DEVICE_CLASS_EDGE_AI
	DeviceClassStorage     DeviceClass = DeviceClass_DEVICE_CLASS_STORAGE
	DeviceClassCompute     DeviceClass = DeviceClass_DEVICE_CLASS_COMPUTE
	DeviceClassSensor      DeviceClass = DeviceClass_DEVICE_CLASS_SENSOR
	DeviceClassRelay       DeviceClass = DeviceClass_DEVICE_CLASS_RELAY
)

const (
	JobKindUnspecified        JobKind = JobKind_JOB_KIND_UNSPECIFIED
	JobKindLlmInference       JobKind = JobKind_JOB_KIND_LLM_INFERENCE
	JobKindImageGeneration    JobKind = JobKind_JOB_KIND_IMAGE_GENERATION
	JobKindVideoTranscode     JobKind = JobKind_JOB_KIND_VIDEO_TRANSCODE
	JobKindDataProcessing     JobKind = JobKind_JOB_KIND_DATA_PROCESSING
	JobKindModelTraining      JobKind = JobKind_JOB_KIND_MODEL_TRAINING
	JobKindEmbedding          JobKind = JobKind_JOB_KIND_EMBEDDING
	JobKindIotAggregation     JobKind = JobKind_JOB_KIND_IOT_AGGREGATION
	JobKindStorageReplication JobKind = JobKind_JOB_KIND_STORAGE_REPLICATION
	JobKindAgentWorkflow      JobKind = JobKind_JOB_KIND_AGENT_WORKFLOW
	JobKindSpatialRecon       JobKind = JobKind_JOB_KIND_SPATIAL_RECON
	JobKindPointcloudAlign    JobKind = JobKind_JOB_KIND_POINTCLOUD_ALIGN
	JobKindMeshOptimize       JobKind = JobKind_JOB_KIND_MESH_OPTIMIZE
	JobKindNerfTrain          JobKind = JobKind_JOB_KIND_NERF_TRAIN
	JobKindSceneRender        JobKind = JobKind_JOB_KIND_SCENE_RENDER
	JobKindCaptureQC          JobKind = JobKind_JOB_KIND_CAPTURE_QC
	JobKindEpochDiff          JobKind = JobKind_JOB_KIND_EPOCH_DIFF
	JobKindExportPack         JobKind = JobKind_JOB_KIND_EXPORT_PACK
	JobKindTranscription      JobKind = JobKind_JOB_KIND_TRANSCRIPTION
)
