package hospitalstay;

import java.util.Scanner;

public class OvernightStay {

    public OvernightStay(){

    }

    public void calculateOvernightCosts(PatientClass patient, Scanner input){
        System.out.println("Did the patient stay overnight? Y or N");

        String yesNo = input.next().toLowerCase();
        if (yesNo.equals("y")){
            boolean overnight;
            do{
                overnight = false;
                System.out.println("Please list the cost for the overnight stay: ");
                double overnightCost = input.nextDouble();
                patient.overnightStayCosts(patient.overnightStayCosts() + overnightCost);
                System.out.println("Is there another overnight stay cost? Y or N : ");
                String overnightContinue = input.next().toLowerCase();
                input.nextLine();
                if (overnightContinue.equals("y")){
                    overnight = true;
                }
            } while(overnight);
        }
    }
}
