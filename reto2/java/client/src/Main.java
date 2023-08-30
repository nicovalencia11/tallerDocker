import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.URL;
import java.net.URLConnection;

public class Main {
    public static void main(String[] args) {
        String usuario = System.getenv().get("USUARIO");
        String serverUrl = System.getenv().get("SERVER_URL");

        try {
            // Esperar 4 segundos (4000 milisegundos)
            Thread.sleep(4000);

            URL url = new URL("http://" + serverUrl + ":80?usuario=" + usuario
                    + "&correo=dobby@gmail.com");
            URLConnection con = url.openConnection();
            BufferedReader resultado = new BufferedReader(new InputStreamReader(con.getInputStream()));
            System.out.println(resultado.readLine());
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            System.out.println("Thread interrupted");
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }
}

