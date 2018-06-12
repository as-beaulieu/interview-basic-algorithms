package hospitalstay;

/*
    Calculate patient charges
        Determine if it is an overnight stay
            Use the overnight charges, medication charges, and labs charges
        If an outpatient stay, then only medication charges and labs charges
    Print the final total
    Then ask the user if they want to calculate for another patient
 */

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
	// write your code here
        Scanner input = new Scanner(System.in);
        OvernightStay overnightStay = new OvernightStay();
        Medications medications = new Medications();
        LabWork labWork = new LabWork();
        CostSummary summary = new CostSummary();

        System.out.println("Welcome to Patient Cost Calculator!");
        boolean calculateContinue;
        do{
            calculateContinue = false;
            PatientClass patient = new PatientClass();

            overnightStay.calculateOvernightCosts(patient, input);
            medications.calculateMedications(patient, input);
            labWork.calculateLabWork(patient, input);
            summary.calculateCostSummary(patient);

            System.out.println("Would you like to calculate costs for another patient? Y or N: ");
            String yesNo = input.next().toLowerCase();
            input.nextLine();
            if (yesNo.equals("y")){
                calculateContinue = true;
            }
        }while(calculateContinue);
    }
}
