import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.URL;
import java.net.URLConnection;

public class Main {
    public static void main(String[] args) {
        String usuario = System.getenv().get("USUARIO");
        String serverUrl = System.getenv().get("SERVER_URL"); // Obtener la URL del servidor desde la variable de entorno

        if (usuario != null && serverUrl != null) {
            try {
                URL url = new URL("http://" + serverUrl + ":80?usuario=" + usuario
                        + "&correo=dobby@gmail.com");
                URLConnection con = url.openConnection();
                BufferedReader resultado = new BufferedReader(new InputStreamReader(con.getInputStream()));
                System.out.println(resultado.readLine());
            } catch (Exception e) {
                System.out.println(e.getMessage());
            }
        } else {
            System.out.println("Falta definir USUARIO y SERVER_URL en las variables de entorno.");
        }
    }
}
