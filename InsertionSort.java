public class InsertionSort {
	
	public static void main(String[] args){
		int[] myArray = new int[]{10,-3,7,2,5,9,4,6};
		
		System.out.println("Initial array.");
		for(int i = 0; i < myArray.length; i++){
			System.out.println(myArray[i]);
		}
		System.out.println();
		
		for(int firstUnsortedIndex = 1; firstUnsortedIndex < myArray.length; firstUnsortedIndex++){
			int newElement = myArray[firstUnsortedIndex];
			int i;
			
			for(i = firstUnsortedIndex; i > 0 && myArray[i-1] > newElement; i--){
				myArray[i] = myArray[i-1];
			}
			myArray[i] = newElement;
		}
		
		System.out.println("Sorted array.");
		for(int i = 0; i < myArray.length; i++){
			System.out.println(myArray[i]);
		}
		
	}
	
}
