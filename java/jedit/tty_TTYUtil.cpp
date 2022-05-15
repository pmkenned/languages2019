// https://www.cse.psu.edu/~kxc104/class/cse472/09s/hw/hw7/vt100ansi.htm
// https://jonisalonen.com/2012/calling-c-from-java-is-easy/
// https://www.baeldung.com/jni
// https://www.infoworld.com/article/2077554/java-tip-54--returning-data-in-reference-arguments-via-jni.html
// https://docs.oracle.com/javase/tutorial/java/nutsandbolts/arrays.html
// https://docs.oracle.com/javase/7/docs/technotes/guides/jni/spec/functions.html
// https://docs.oracle.com/javase/1.5.0/docs/guide/jni/spec/types.html
// https://docs.oracle.com/javase/8/docs/technotes/guides/jni/spec/design.html
// https://stackoverflow.com/questions/24641536/how-to-set-java-home-in-linux-for-all-users

#include "tty_TTYUtil.h"
#include <stdlib.h>
#include <termios.h>
#include <unistd.h>

struct termios orig_termios;

static void
die(const char *s) {
    perror(s);
    exit(1);
}

JNIEXPORT void JNICALL Java_tty_TTYUtil_disableRawMode
  (JNIEnv * env, jclass cls)
{
  if (tcsetattr(STDIN_FILENO, TCSAFLUSH, &orig_termios) == -1)
    die("tcsetattr");
}

static void disableRawMode()
{
    Java_tty_TTYUtil_disableRawMode(NULL, NULL);
}

JNIEXPORT void JNICALL Java_tty_TTYUtil_enableRawMode
  (JNIEnv * env, jclass cls)
{
  if (tcgetattr(STDIN_FILENO, &orig_termios) == -1) {
      die("tcgetattr");
  }
  atexit(disableRawMode);
  struct termios raw = orig_termios;
  raw.c_iflag &= ~(BRKINT | ICRNL | INPCK | ISTRIP | IXON);
  raw.c_oflag &= ~(OPOST);
  raw.c_cflag |= (CS8);
  raw.c_lflag &= ~(ECHO | ICANON | IEXTEN | ISIG);
  raw.c_cc[VMIN] = 0;
  raw.c_cc[VTIME] = 1;
  if (tcsetattr(STDIN_FILENO, TCSAFLUSH, &raw) == -1) {
      die("tcsetattr");
  }
}

JNIEXPORT jstring JNICALL Java_tty_TTYUtil_getTTYName
  (JNIEnv *env, jclass cls)
{
    char *name = ttyname(STDOUT_FILENO);
    return env->NewStringUTF(name);
}

JNIEXPORT jboolean JNICALL Java_tty_TTYUtil_isTTY
  (JNIEnv *env, jclass cls)
{
    return isatty(STDOUT_FILENO)? JNI_TRUE: JNI_FALSE;
}

JNIEXPORT jint JNICALL Java_tty_TTYUtil_tcgetattr
  (JNIEnv * env, jclass cls, jint fd, jobject termios)
{
    struct termios t;
    jint result = tcgetattr(fd, &t);

    jclass termiosClass = env->GetObjectClass(termios);
    jfieldID fid;

    fid = env->GetFieldID(termiosClass, "IFlag", "I");
    env->SetIntField(termios, fid, t.c_iflag);

    fid = env->GetFieldID(termiosClass, "OFlag", "I");
    env->SetIntField(termios, fid, t.c_oflag);

    fid = env->GetFieldID(termiosClass, "CFlag", "I");
    env->SetIntField(termios, fid, t.c_cflag);

    fid = env->GetFieldID(termiosClass, "LFlag", "I");
    env->SetIntField(termios, fid, t.c_lflag);

    // https://stackoverflow.com/questions/1086596/how-to-access-arrays-within-an-object-with-jni
    //fid = env->GetFieldID(termiosClass, "CC", "[I");
    //env->SetIntField(termios, fid, t.c_cc);

    return result;
}

JNIEXPORT jint JNICALL Java_tty_TTYUtil_tcsetattr
  (JNIEnv * env, jclass cls, jint fd, jint optional_actions, jobject termios)
{
    struct termios t;

    jclass termiosClass = env->GetObjectClass(termios);
    jfieldID fid;

    fid = env->GetFieldID(termiosClass, "IFlag", "I");
    t.c_iflag = env->GetIntField(termios, fid);

    fid = env->GetFieldID(termiosClass, "OFlag", "I");
    t.c_oflag = env->GetIntField(termios, fid);

    fid = env->GetFieldID(termiosClass, "CFlag", "I");
    t.c_cflag = env->GetIntField(termios, fid);

    fid = env->GetFieldID(termiosClass, "LFlag", "I");
    t.c_lflag = env->GetIntField(termios, fid);

    // fid = env->GetFieldID(termiosClass, "CC", "[I");
    // t.c_cc = env->GetIntField(termios, fid);

    jint result = tcsetattr(fd, optional_actions, &t);

    return result;
}
