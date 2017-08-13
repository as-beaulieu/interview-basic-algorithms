using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

//Always best to name your files after your class for better organization

//Static in c#
/*
class SomeClass {
	public int InstanceMethod() { return 1; }
	public static int StaticMethod() { return 42; }
}
SomeClass instance = new SomeClass();

	instance.InstanceMethod();	//Fine
	instance.StaticMethod();	//Won't Compile!

	SomeClass.InstanceMethod();	//Won't Compile!
	SomeClass.StaticMethod();	//Fine

With a static method you do not need and instance of the class to call the method, just the class
*/

//namespace segregates code, such as Player
namespace guessingGame
{
    //Classes can be given instances with their own variables of given 'characteristics'
    class Program 
    {
        //() notes a function, but inside () goes the parameters the function intakes
        //Void is what the function outputs. Void means the () returns nothing
        //Main() is required for the program to compile, running from the Main()
        //'static' () means not associated with an instance of the class
        static void Main(string[] args)
        {
            Console.WriteLine("Let's play a guessing game!");
            Console.WriteLine("Guess a number between 1 and 10...");
            Console.WriteLine("What is your guess?");

            int answer = computerPicks(); //This what the computer is thinking
            int guessInt;
            
            //Game loop, will allow game to continue until condition is met
            do
            {
                guessInt = guess();
                if (guessInt == -1)
                {
                    endGame();
                    return;
                }
                //Handles the comparison between answer and guessInt
                if (answer == guessInt)
                {
                    Console.WriteLine("Correct! I was thinking " + answer);
                    //keepGoing = false;
                    //Another way is to use 'break' which will exit the most current loop
                }
                else if (answer > guessInt)
                {
                    Console.WriteLine("Higher! Try Again.");
                }
                else //Last only option is answer < guess
                {
                    Console.WriteLine("Lower! Try Again.");
                }
            }
            while (answer != guessInt);

            endGame();
            //"Returns" out of current function, which is Main(), so ends program
            return;
        }
        //End Main()

        //guess() handles user input
        public static int guess()
        {
            //Ask the user for their guess with ReadLine()
            string guess = Console.ReadLine();
            Console.WriteLine("You guessed: " + guess);

            //Making sure that user inputs a int
            try
            {
                //Converting guess string to Int for >,<,= comparison
                int num = Int32.Parse(guess);
                return num;
            }
            //If the user did not input a Int
            catch (Exception e)
            {
                Console.WriteLine("That is not a number!");
                Console.WriteLine("The error was: " + e);
                //Have to return an int, so -1 to confirm an error occurred.
                return -1;
            }
        }
        //End guess()

        //computerPicks() handles the random # the comp 'thinks' of
        //Note: Pseudorandom, which means 'random enough'
        public static int computerPicks()
        {
            //'Instantiate' the Random class to make our own random generator for this instance
            //and keeping a copy of this instance in a variable 'random'
            Random randomNum = new Random();
            //The second int is exclusive, meaning you have to go +1 of # desired
            int numbers = randomNum.Next(1, 11);
            return numbers;
        }
        //End computerPicks()

        //endGame() handles the assurance that the console stays up so the user
        //can read any final messages, good or bad.
        public static void endGame()
        {
            //End of game
            Console.WriteLine("Press any key to continue...");
            Console.ReadLine();
        }
        //End endGame()
    }
    //end Class Program
}
//end namespace guessingGame
