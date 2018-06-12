package hospitalstay;

public class CostSummary {

    public CostSummary(){

    }

    public void calculateCostSummary(PatientClass patient){
        double totalCosts = patient.overnightStayCosts() + patient.medicationCosts() + patient.labsCosts();
        System.out.println("The total cost for overnight stay for this patient is: " + patient.overnightStayCosts());
        System.out.println("The total cost for medications for this patient is: " + patient.medicationCosts());
        System.out.println("The total cost for lab work for this patient is: " + patient.labsCosts());
        System.out.println("The total overall cost for this patient is: " + totalCosts);
    }
}
