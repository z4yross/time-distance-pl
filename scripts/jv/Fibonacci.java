package scripts.jv;

import java.text.NumberFormat;
import java.text.DecimalFormat;

class Main{

    public static void main(String args[]){
        long inicioF = System.currentTimeMillis();
        int n=1, f=1;
        for(double i=0; i<10000000; i++){
            int t = f;
            f += n;
            n=t;
        }
        long finF = System.currentTimeMillis();
        NumberFormat formatter = new DecimalFormat("#0.0000");
        double tiempoF = ((double) (finF-inicioF) / 1000.0);
        System.out.print(formatter.format(tiempoF));
    }

}