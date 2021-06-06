public class MergeSortAlgorithm{
	public static void main(String[] args){
		int[] myArray = new int[]{99,48,60,-17,51,10,13};
		
		System.out.println("Original items");
		for(int i = 0; i < myArray.length; i++){
			System.out.println(myArray[i]);
		}
		
		System.out.println();
		
		mergeSort(myArray, 0, myArray.length);
		
		System.out.println("Sorted items");
		for(int i = 0; i < myArray.length; i++){
			System.out.println(myArray[i]);
		}
	}
	
	public static void mergeSort(int[] input, int start, int end){
		//{20,35,-15,7,55,1,-22};
		if((end - start) < 2){
			return;
		}
		
		//get mid for the left side of the split array
		int mid = (start + end)/2;
		mergeSort(input, start, mid);
		mergeSort(input, mid, end);
		merge(input, start, mid, end);
	}
	
	public static void merge(int[] input, int start, int mid, int end){
		//{20,35,-15,7,55,1,-22};
		if(input[mid-1] <= input[mid]){
			return;
		}
		
		int i = start;
		int j = mid;
		int tempIndex = 0;
		
		int[] temp = new int[end -start];
		while(i < mid && j < end){
			temp[tempIndex++] = input[i] <= input[j] ? input[i++] : input[j++];
		}
		
		System.arraycopy(input, i, input, start + tempIndex, mid - i);
		System.arraycopy(temp, 0, input, start, tempIndex);
	}
}