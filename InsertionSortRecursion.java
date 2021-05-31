public class InsertionSortRecursion {
	
	public static void main(String[] args){
		int[] myArray = new int[]{10,-3,7,2,5,9,4,6};
		
		System.out.println("Initial array.");
		for(int i = 0; i < myArray.length; i++){
			System.out.println(myArray[i]);
		}
		System.out.println();
		
		insertionSort(myArray, myArray.length);
		
		// for(int firstUnsortedIndex = 1; firstUnsortedIndex < myArray.length; firstUnsortedIndex++){
			// int newElement = myArray[firstUnsortedIndex];
			// int i;
			
			// for(i = firstUnsortedIndex; i > 0 && myArray[i-1] > newElement; i--){
				// myArray[i] = myArray[i-1];
			// }
			// myArray[i] = newElement;
		// }
		
		System.out.println("Sorted array.");
		for(int i = 0; i < myArray.length; i++){
			System.out.println(myArray[i]);
		}
		
	}
	
	public static void insertionSort(int[] input, int numItems){
		if (numItems < 2)
			return;
		
		insertionSort(input, numItems - 1);
		int newElement = input[numItems - 1];
		int i;
		
		for(i = numItems -1; i > 0 && input[i-1] > newElement; i--){
			input[i] = input[i-1];
		}
		
		input[i] = newElement;
		
		System.out.println("Result of call when numItems = " + numItems);
		for(i = 0; i < input.length; i++){
			System.out.print(input[i]);
			System.out.print(", ");
		}
		System.out.print("");
		System.out.println("-----------------------");
	}
}
