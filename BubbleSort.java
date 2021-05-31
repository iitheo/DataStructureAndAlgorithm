public class BubbleSort {
	public static void main(String[] args){
		int[] myArray = new int[]{10,-3,7,2,5,9,4,6};
		System.out.println("Initial array.");
		for(int i = 0; i < myArray.length; i++){
			System.out.println(myArray[i]);
		}
		for (int lastUunsortedIndex = myArray.length -1; lastUunsortedIndex > 0; lastUunsortedIndex--){
			for (int i = 0; i < lastUunsortedIndex; i++){
				if (myArray[i] > myArray[i+1]){
					swap(myArray, i, i+1);
				}
			}
		}
		
		System.out.println("After sorting array.");
		for(int i = 0; i < myArray.length; i++){
			System.out.println(myArray[i]);
		}
	}
	
	public static void swap(int[] myArray, int i, int j){
		if(i == j){
			return;
		}
		
		int temp = myArray[i];
		myArray[i] = myArray[j];
		myArray[j] = temp;
	}
}
