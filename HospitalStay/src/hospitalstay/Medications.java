package hospitalstay;

import java.util.Scanner;

public class Medications {

    public Medications(){

    }

    public void calculateMedications(PatientClass patient, Scanner input){
        System.out.println("Did the patient need any medications? Y or N :");

        String yesNo = input.next().toLowerCase();
        if (yesNo.equals("y")){
            boolean meds;
            do{
                meds = false;
                System.out.println("Please list the cost for the patient's medication: ");
                double medCost = input.nextDouble();
                patient.medicationCosts(patient.medicationCosts() + medCost);
                System.out.println("Is there another medication cost? Y or N :");
                yesNo = input.next().toLowerCase();
                input.nextLine();
                if (yesNo.equals("y")){
                    meds = true;
                }
            } while(meds);
        }
    }
}
