package flow

import (
	"fmt"
)

func (fc *FlowContext) OnInterrupt() {

	fmt.Print("\n")
	for _, step := range fc.Steps {
		if step.Output != nil {
			fmt.Printf("step:%s%d\n", step.Name, step.Id)
			for _, input := range step.Inputs {
				fmt.Printf("  input  : d%d\n", input.Id)
				for _, shard := range input.Shards {
					printShardStatus(shard)
				}
			}
			fmt.Printf("  output : d%d\n", step.Output.Id)
			for _, task := range step.Tasks {
				for _, shard := range task.Outputs {
					printShardStatus(shard)
				}
			}
		}
	}
	fmt.Print("\n")
}

func printShardStatus(shard *DatasetShard) {
	if shard.Closed() {
		fmt.Printf("     shard:%d time:%v completed %d\n", shard.Id, shard.TimeTaken(), shard.Counter)
	} else {
		fmt.Printf("     shard:%d time:%v processed %d\n", shard.Id, shard.TimeTaken(), shard.Counter)
	}
}
