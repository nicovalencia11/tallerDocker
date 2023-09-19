package com.example.autenticacion.services.implementation;

import com.example.autenticacion.entities.Usuario;
import com.example.autenticacion.repositories.UsuarioRepository;
import com.example.autenticacion.services.services.UsuarioService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UsuarioServiceImpl implements UsuarioService {

    @Autowired
    private UsuarioRepository usuarioRepository;

    /**
     * Metodo que permite registrar un usuario
     *
     * @param usuario
     * @return
     */
    @Override
    public Usuario registrarUsuario(Usuario usuario) throws Exception {
        Usuario usuarioSave = usuarioRepository.save(usuario);
        if(usuarioSave == null){
            throw new Exception("Error en el registro del usuario");
        }
        return usuarioSave;
    }

    /**
     * Metodo que permite el login con nombre de usuario y contraseña
     *
     * @param nombreUsuario
     * @param password
     * @return
     */
    @Override
    public Usuario loginUsuario(String nombreUsuario, String password) throws Exception {
        Usuario usuario = usuarioRepository.findByNombreUsuarioAndPassword(nombreUsuario,password);
        if (usuario == null){
            throw new Exception("Los datos de autenticacion no son correctos");
        }
        return usuario;
    }

    /**
     * Metodo que permite actualizar un usuario
     *
     * @param usuario
     * @return
     */
    @Override
    public Usuario actualizarUsuario(Usuario usuario) throws Exception {
        Usuario usuarioSave = usuarioRepository.save(usuario);
        if(usuarioSave == null){
            throw new Exception("Error en la actualizacion del usuario");
        }
        return usuarioSave;
    }

    /**
     * metodo que permite listar todos los usuarios
     *
     * @return
     */
    @Override
    public List<Usuario> listarUsuarios() {
        return usuarioRepository.findAll();
    }

    /**
     * Metodo que permite recuperar la clave un usuario
     *
     * @param correo
     * @return
     */
    @Override
    public Usuario recuperarPassword(String correo) throws Exception {
        Usuario usuarioSave = usuarioRepository.findByCorreo(correo);
        if(usuarioSave == null){
            throw new Exception("Error, el correo no fue encontrado");
        }
        return usuarioSave;
    }

}
