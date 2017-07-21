
/**
 * FizzBuzz is a common interview question for programmers.
 * 
 * @author (your name) 
 * @version (a version number or a date)
 */
import java.util.*;
//Arrays are built in, but of fixed size. Importing is not neccessary.
public class FizzBuzz{
    public static void main(String[] args) {
        for (int i = 1; i <= 100; i++) {
            boolean fizzOrBuzz = false;
            if (i % 3 == 0) {
                System.out.print("Fizz");
                fizzOrBuzz = true;
            }
            if (i % 5 == 0) {
                System.out.print("Buzz");
                fizzOrBuzz = true;
            }

            if (fizzOrBuzz) {
                System.out.println();
            } else {
                System.out.println(String.valueOf(i));
            }
        }
    }

}
