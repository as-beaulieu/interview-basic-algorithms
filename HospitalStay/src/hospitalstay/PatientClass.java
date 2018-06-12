package hospitalstay;

public class PatientClass {

    private double overnightStayCosts = 0.0;

    public double overnightStayCosts() { return overnightStayCosts; } //getter
    public void overnightStayCosts(double value) { overnightStayCosts = value; } //setter

    private double medicationCosts = 0.0;

    public double medicationCosts() { return medicationCosts; }
    public void medicationCosts(double value) { medicationCosts = value; }

    private double labsCosts = 0.0;

    public double labsCosts() { return labsCosts; }
    public void labsCosts(double value) { labsCosts = value; }
}
