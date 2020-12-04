package perfstat

type PartitionType struct {
	bool SmtCapable        /* OS supports SMT mode */
	bool SmtEnabled        /* SMT mode is on */
	bool LparCapable       /* OS supports logical partitioning */
	bool LparEnabled       /* logical partitioning is on */
	bool SharedCapable     /* OS supports shared processor LPAR */
	bool SharedEnabled     /* partition runs in shared mode */
	bool DLparCapable      /* OS supports dynamic LPAR */
	bool Capped            /* partition is capped */
	bool Kernel64bit       /* kernel is 64 bit */
	bool PoolUtilAuthority /* pool utilization available */
	bool DonateCapable     /* capable of donating cycles */
	bool DonateEnabled     /* enabled for donating cycles */
	bool AmsCapable        /* 1 = AMS(Active Memory Sharing) capable, 0 = Not AMS capable */
	bool AmsEnabled        /* 1 = AMS(Active Memory Sharing) enabled, 0 = Not AMS enabled */
	bool PowerSave         /*1= Power saving mode is enabled*/
	bool AmeEnabled        /* Active Memory Expansion is enabled */
	bool SharedExtended
}

type PartitionValue struct {
	int64 Online
	int64 Max
	int64 Min
	int64 Desired
}

type PartitionConfig struct {
	int64          Version                  /* Version number */
	string         Name                     /* Partition Name */
	string         Node                     /* Node Name */
	PartitionType  Conf                     /* Partition Properties */
	int32          Number                   /* Partition Number */
	int32          GroupID                  /* Group ID */
	string         ProcessorFamily          /* Processor Type */
	string         ProcessorModel           /* Processor Model */
	string         MachineID                /* Machine ID */
	float64        ProcessorMhz             /* Processor Clock Speed in MHz */
	PartitionValue NumProcessors            /* Number of Configured Physical Processors in frame*/
	string         OSName                   /* Name of Operating System */
	string         OSVersion                /* Version of operating System */
	string         OSBuild                  /* Build of Operating System */
	int32          LCpus                    /* Number of Logical CPUs */
	int32          SmtThreads               /* Number of SMT Threads */
	int32          Drives                   /* Total Number of Drives */
	int32          NetworkAdapters          /* Total Number of Network Adapters */
	PartitionValue CpuCap                   /* Min, Max and Online CPU Capacity */
	int32          Weightage                /* Variable Processor Capacity Weightage */
	int32          EntCapacity              /* number of processor units this partition is entitled to receive */
	PartitionValue VCpus                    /* Min, Max and Online Virtual CPUs */
	int32          PoolID                   /* Shared Pool ID of physical processors, to which this partition belongs*/
	int32          ActiveCpusInPool         /* Count of physical CPUs in the shared processor pool, to which this partition belongs */
	int32          PoolWeightage            /* Pool Weightage */
	int32          SharedPCpu               /* Number of physical processors allocated for shared processor use */
	int32          MaxPoolCap               /* Maximum processor capacity of partition's pool */
	int32          EntPoolCap               /* Entitled processor capacity of partition's pool */
	PartitionValue Mem                      /* Min, Max and Online Memory */
	int32          MemWeightage             /* Variable Memory Capacity Weightage */
	int64          TotalIOMemoryEntitlement /* I/O Memory Entitlement of the partition in bytes */
	int32          MemPoolID                /* AMS pool id of the pool the LPAR belongs to */
	int64          HyperPgSize              /* Hypervisor page size in KB*/
	PartitionValue ExpMem                   /* Min, Max and Online Expanded Memory */
	int64          TargetMemExpFactor       /* Target Memory Expansion Factor scaled by 100 */
	int64          TargetMemExpSize         /* Expanded Memory Size in MB */
	int32          SubProcessorMode         /* Split core mode, its value can be 0,1,2 or 4. 0 for unsupported, 1 for capable but not enabled, 2 or 4 for enabled*/
}
