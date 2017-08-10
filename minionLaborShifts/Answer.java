/**
Commander Lamda's minions are upset! They're given the worst jobs in the whole
space station, and some of them are starting to complain that even those worst
jobs are being allocated unfairly. If you can fix this problem, it'll prove
your chops to Commander Lambda so you can get promoted!

Minion's tasks are assigned by putting their ID numbers into a list, one time
for each day they'll work that task. As shifts are planned well in advance, the
lists of each task will contain up to 99 integers. When a minion is scheduled 
for the same task too many times, they'll complain about it until they're taken
off the task completely. Some tasks are worse that others, so the number of 
scheduled assignments beforea minion will refuse to do a task varies depending on 
the task. You figure you can speed things up by automating the removal of the
minions who have been assigned a task too many times before they even get a 
chance to start complaining.

----------------------------------------------------------------------

Write a function called answer(data, n) that takes in a list of less than
100 integers and a number n, and returns that same list but with all of the 
numbers that occur more than n times removed entirely. The returned list should
retain the same ordering as the original list - you don't want to mix up those
carefully-planned shift rotations! For instance, if data was [5, 10, 15, 10, 7]
and n = 1, answer(data, n) would return the list [5, 15, 7] because 10 occurs
twice, and thus was removed from the list entirely.

Test Cases
----------
Input:
    (int list) data = [1,2,3]
    (int) n = 0
Output: (int list) []

Input: 
    (int list) data = [1,2,2,3,3,3,4,5,5]
    (int) n = 1
Output: (int list) [1,4]

Input: 
    (int list) data = [1,2,3]
    (int) n = 6
Output: (int list) [1,2,3]
**/
 
//package com.google.challenges;
 
import java.util.*;

public class Answer
{
    public Answer(int[] data, int n) {
        
        
        
        System.out.println(Arrays.toString(data));
        System.out.println(n);
        
        
        
    }
    
    public static int[] trim(int[] data, int n) {
        System.out.println(Arrays.toString(data));
        System.out.println("Old array length: " + data.length);
        int deleteCount = 0;
        for (int i = 0; i < data.length; i++) {
            int numDup = 0;
            int check = data[i];
            for (int j = 0; j < data.length; j++) {
                if (check == data[j] && data[j] != 0) {
                    numDup++;
                }
                if (numDup > n) {
                    int delCheck = data[j];
                    for (int k = 0; k < data.length; k++) {
                        if (delCheck == data[k] && data[k] != 0){
                            System.out.println(delCheck + " exceeds # of duplications!");
                            data[k] = 0;
                            deleteCount++;
                        }
                    }
                }
            }
        }
        
        System.out.println("Number of deletions: " + deleteCount);        

        int[] a = new int[data.length - deleteCount];
        System.out.println("New Array size: " + a.length);
        int l = 0;
        for (int i = 0; i < data.length; i++) {
            if (data[i] != 0) {
                a[l] = data[i];
                l++;
            }
        }
        return a;
    }
    
}
