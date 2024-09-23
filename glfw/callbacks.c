#include "GLFW/glfw3.h"

void goCursorPosCallback(GLFWwindow *window, double x, double y);
void goSizeCallback(GLFWwindow *window, int width, int height);
void goFramebufferSizeCallback(GLFWwindow *window, int width, int height);

void glfwSetCursorPosCallback_fix(GLFWwindow *window) {
    glfwSetCursorPosCallback(window, (GLFWcursorposfun)goCursorPosCallback);
}

void glfwSetWindowSizeCallback_fix(GLFWwindow *window) {
    glfwSetWindowSizeCallback(window, (GLFWwindowsizefun)goSizeCallback);
}

void glfwSetFramebufferSizeCallback_fix(GLFWwindow *window) {
    glfwSetFramebufferSizeCallback(window, (GLFWframebuffersizefun)goFramebufferSizeCallback);
}

