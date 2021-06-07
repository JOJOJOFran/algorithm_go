package array

func MoveZeroes(nums []int){
	//找到非0需要填充到位置
	
	//for i:=0;i< len(nums);i++{
	//	if nums[i]!=0{
	//		tmp:= nums[j]
	//		nums[j] = nums[i]
	//		nums[i] = tmp
	//		j++
	//	}
	//}
	
	for i,j:=0,0;i< len(nums);i++{
		if nums[i]!=0{
			if(i!=j){
				nums[j]=nums[i]
				nums[i]=0
			}
			j++
		}
	}
}