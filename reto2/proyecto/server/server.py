from http.server import BaseHTTPRequestHandler, HTTPServer
import urllib.parse

class MyServer(BaseHTTPRequestHandler):
    def do_GET(self):
        parsed_path = urllib.parse.urlparse(self.path)
        query = urllib.parse.parse_qs(parsed_path.query)
        if 'usuario' in query:
            usuario = query['usuario'][0]
            self.send_response(200)
            self.send_header('Content-type', 'text/html')
            self.end_headers()
            self.wfile.write(bytes(f'Hola {usuario}', 'utf-8'))
        else:
            self.send_response(400)
            self.send_header('Content-type', 'text/html')
            self.end_headers()
            self.wfile.write(bytes('Bad Request', 'utf-8'))

if __name__ == '__main__':
    server = HTTPServer(('0.0.0.0', 80), MyServer)
    server.serve_forever()
