#include <cstdlib>
#include <iostream>

#include <GL/glew.h>
#include <SDL2/SDL.h>

GLuint program;
GLuint attribute_coord2d;

bool init_resources(void) {
    GLint compile_ok = GL_FALSE, link_ok = GL_FALSE;

    GLuint vs = glCreateShader(GL_VERTEX_SHADER);
    const char * vs_source = 
        "#version 120\n"
        "attribute vec2 coord2d;\n"
        "void main(void) {\n"
        "    gl_Position = vec4(coord2d, 0.0, 1.0);\n"
        "}\n";
    glShaderSource(vs, 1, &vs_source, NULL);
    glCompileShader(vs);
    glGetShaderiv(vs, GL_COMPILE_STATUS, &compile_ok);
    if(!compile_ok) {
        std::cerr << "Error in vertex shader" << std::endl;
        return false;
    }

    GLuint fs = glCreateShader(GL_FRAGMENT_SHADER);
    const char * fs_source =
        "#version 120\n"
        "void main(void) {\n"
        "    gl_FragColor[0] = 0.0;\n"
        "    gl_FragColor[1] = 0.0;\n"
        "    gl_FragColor[2] = 1.0;\n"
        "}";
    glShaderSource(fs, 1, &fs_source, NULL);
    glCompileShader(fs);
    glGetShaderiv(fs, GL_COMPILE_STATUS, &compile_ok);
    if(!compile_ok) {
        std::cerr << "Error in fragment shader" << std::endl;
        return false;
    }

    program = glCreateProgram();
    glAttachShader(program, vs);
    glAttachShader(program, fs);
    glLinkProgram(program);
    glGetProgramiv(program, GL_LINK_STATUS, &link_ok);
    if(!link_ok) {
        std::cerr << "Error in glLinkProgram" << std::endl;
        return false;
    }

    const char * attribute_name = "coord2d";
    attribute_coord2d = glGetAttribLocation(program, attribute_name);
    if(attribute_coord2d == -1) {
        std::cerr << "Could not bind attribute" << std::endl;
        return false;
    }

    return true;
}

void render(SDL_Window * window) {
    glClearColor(1.0, 1.0, 1.0, 1.0);
    glClear(GL_COLOR_BUFFER_BIT);

    glUseProgram(program);
    glEnableVertexAttribArray(attribute_coord2d);
    GLfloat triangle_vertices[] = {
         0.0,  0.8,
        -0.8, -0.8,
         0.8, -0.8,
    };
    glVertexAttribPointer(
        attribute_coord2d,
        2,
        GL_FLOAT,
        GL_FALSE,
        0,
        triangle_vertices
    );

    glDrawArrays(GL_TRIANGLES, 0, 3);
    glDisableVertexAttribArray(attribute_coord2d);
    SDL_GL_SwapWindow(window);
}

void free_resources() {
    glDeleteProgram(program);
}

void mainLoop(SDL_Window * window) {
    while(true) {
        SDL_Event ev;
        while (SDL_PollEvent(&ev)) {
            if(ev.type == SDL_QUIT) {
                return;
            }
        }
        render(window);
    }
}

int main(int argc, char * argv[]) {
    SDL_Init(SDL_INIT_VIDEO);
    SDL_Window * window = SDL_CreateWindow("My First Triangle",
                                           SDL_WINDOWPOS_CENTERED,
                                           SDL_WINDOWPOS_CENTERED,
                                           640, 480,
                                           SDL_WINDOW_RESIZABLE | SDL_WINDOW_OPENGL);
    SDL_GL_CreateContext(window);

    GLenum glew_status = glewInit();
    if(glew_status != GLEW_OK) {
        std::cerr << "Error: glewInit: " << glewGetErrorString(glew_status) << std::endl;
        return EXIT_FAILURE;
    }

    if(!init_resources())
        return EXIT_FAILURE;

    mainLoop(window);

    free_resources();
    return EXIT_SUCCESS;
}
