public class ShellSort{
	public static void main(String[] args){
		int[] myArray = new int[]{99,48,60,-17,51,10,13};
		System.out.println("Original items");
		for(int i = 0; i < myArray.length; i++){
			System.out.println(myArray[i]);
		}
		System.out.println();
		
		for(int gap = myArray.length/2; gap > 0; gap /= 2){
			for(int i = gap; i < myArray.length; i++){
				int newElement = myArray[i];
				int j = i;
				while(j >= gap && myArray[j-gap] > newElement){
					myArray[j] = myArray[j-gap];
					j -= gap;
				}
				myArray[j] = newElement;
			}
		}
		System.out.println();
		
		System.out.println("Sorted items");
		for(int i = 0; i < myArray.length; i++){
			System.out.println(myArray[i]);
		}
		System.out.println();
	}
}