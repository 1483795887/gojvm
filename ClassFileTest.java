package jvmgo.ch05;

public class ClassFileTest {
  public static final boolean FLAG = true;
  public static final byte BYTE = 123;
  public static final char X = 'X';
  public static final short SHORT = 12345;
  public static final int INT = 123456789;
  public static final long LONG = 12345678901L;
  public static final float PI = 3.141f;
  public static final double E = 2.71828;

  public static float circumference(float r){
    float pi = 3.14f;
    float c = 2 * pi * r;
    return c;
  }

  public static void main(String[] args) throws RuntimeException {
    System.out.println("Hello, World!");
  }
}