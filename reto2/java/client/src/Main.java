import java.io.BufferedReader;
import java.net.URL;
import java.net.URLConnection;

public class Main {
    public static void main(String[] args) {
        String usuario = System.getenv().get("USUARIO");
        URL url;
        try {
            Thread.sleep(2000); // Pausa por 2 segundos
        } catch (InterruptedException e) {
            System.out.println(e.getMessage());
        }
        try {
            url = new URL("http://server:80?usuario=" + usuario);
            URLConnection con = url.openConnection();
            BufferedReader resultado = new BufferedReader(new java.io.InputStreamReader(con.getInputStream()));
            System.out.println(resultado.readLine());
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }
}