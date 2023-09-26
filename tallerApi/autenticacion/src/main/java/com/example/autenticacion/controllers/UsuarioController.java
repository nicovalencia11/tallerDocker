package com.example.autenticacion.controllers;

import com.example.autenticacion.entities.Usuario;
import com.example.autenticacion.services.implementation.UsuarioServiceImpl;
import com.example.autenticacion.services.services.UsuarioService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/v1/usuario")
public class UsuarioController {

    @Autowired
    private UsuarioServiceImpl usuarioService;

    /**
     * metodo para registrar un usuario en el sistema
     * @param usuario
     * @return
     * @throws Exception
     */
    @PostMapping("/registrar")
    public ResponseEntity<String> registrar (@RequestBody Usuario usuario) throws Exception{
        usuarioService.registrarUsuario(usuario);
        return ResponseEntity.status(HttpStatus.OK).body("Resgistro Exitoso");
    }

    /**
     * metodo para actualizar un usuario en el sistema
     * @param usuario
     * @return
     * @throws Exception
     */
    @PostMapping("/actualizar")
    public ResponseEntity<String> actualizar (@RequestHeader("token") String token, @RequestBody Usuario usuario) throws Exception{
        Usuario usuario1 = usuarioService.buscarUsuario(usuario);
        usuario1.setPassword(usuario.getPassword());
        if(token.equals(usuario1.getToken())){
            usuarioService.actualizarUsuario(usuario1);
            return ResponseEntity.status(HttpStatus.OK).body("Actualizacion Exitosa");
        }else{
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("No tienes permisos para actualizar este usuario");
        }
    }

    /**
     * metodo para listar los usuarios
     * @param pagina
     * @param tamano
     * @return
     * @throws Exception
     */
    @GetMapping("/listar")
    public Page<Usuario> listar (@RequestParam(name = "pagina", defaultValue = "0") int pagina,
            @RequestParam(name = "tamano", defaultValue = "10") int tamano) throws Exception{
       return usuarioService.listarUsuarios(pagina, tamano);
    }

    /**
     * metodo para recuperar la clave de un usuario en el sistema
     * @param correo
     * @return
     * @throws Exception
     */
    @PostMapping("/recuperarPassword")
    public ResponseEntity<String> recuperarPassword (@RequestBody String correo) throws Exception{
        Usuario usuario = usuarioService.recuperarPassword(correo);
        return ResponseEntity.status(HttpStatus.OK).body("tu clave es: "+usuario.getPassword());
    }
}
