package hospitalstay;

import java.util.Scanner;

public class LabWork {

    public LabWork() {

    }

    public void calculateLabWork(PatientClass patient, Scanner input){
        System.out.println("Did the patient need any lab work? Y or N :");

        String yesNo = input.next().toLowerCase();
        if (yesNo.equals("y")){
            boolean labs;
            do{
                labs = false;
                System.out.println("Please list the cost for the patient's lab work: ");
                double labCost = input.nextDouble();
                patient.labsCosts(patient.labsCosts() + labCost);
                System.out.println("Is there another lab cost? Y or N: ");
                yesNo = input.next().toLowerCase();
                input.nextLine();
                if(yesNo.equals("y")){
                    labs = true;
                }
            } while (labs);
        }
    }
}
