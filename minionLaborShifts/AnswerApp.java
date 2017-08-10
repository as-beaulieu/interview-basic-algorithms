
/**
 * Write a description of class AnswerApp here.
 * 
 * @author (your name) 
 * @version (a version number or a date)
 */

import java.util.Scanner;
import java.util.Arrays;
import javax.swing.JOptionPane;

public class AnswerApp
{
    public static void main(String[] args) {
        int n;
        int[] result;
        
        Scanner kb = new Scanner (System.in);
     
        System.out.println("How many duplicates before you like it removed? 1 - 10?");
        n = kb.nextInt();

        int[] data = {1, 2, 3, 3, 4, 2, 5, 6, 6, 7, 8, 3, 4, 9, 10, 3, 11, 12, 3};
        int[] a = Answer.trim(data, n);
        
        System.out.println("New array: " + Arrays.toString(a));
    }
}
