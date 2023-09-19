package com.example.autenticacion.controllers;

import com.example.autenticacion.entities.Usuario;
import com.example.autenticacion.services.implementation.UsuarioServiceImpl;
import com.example.autenticacion.services.services.UsuarioService;
import org.springframework.beans.factory.annotation.Autowired;
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
    public ResponseEntity<String> actualizar (@RequestBody Usuario usuario) throws Exception{
        usuarioService.actualizarUsuario(usuario);
        //validar token
        return ResponseEntity.status(HttpStatus.OK).body("Actualizacion Exitosa");
    }

    /**
     * metodo para listar los usuarios
     * @return
     * @throws Exception
     */
    @GetMapping("/listar")
    public List<Usuario> listar () throws Exception{
       return usuarioService.listarUsuarios();
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
