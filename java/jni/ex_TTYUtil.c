#include "ex_TTYUtil.h"
#include <unistd.h>

JNIEXPORT jstring JNICALL Java_ex_TTYUtil_getTTYName
  (JNIEnv *env, jclass cls)
{
    char *name = ttyname(STDOUT_FILENO);
    return (*env)->NewStringUTF(env, name);
}

JNIEXPORT jboolean JNICALL Java_ex_TTYUtil_isTTY
  (JNIEnv *env, jclass cls)
{
    return isatty(STDOUT_FILENO)? JNI_TRUE: JNI_FALSE;
}
