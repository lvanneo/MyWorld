package com.example.util;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.RandomAccessFile;

public class Test {
	static {
		System.out.println("静态代码执行！！");
		System.out.println();
	}
	
	public static void main2(String[] args) {
		byte[] bb = {'r','h','m'};
		String ss = new String(bb);
		System.out.println(ss);
	}
	
	public static void main(String[] args) throws Exception {
		Test.split("F:\\Temp\\wpfImage.zip", 1024 * 1024);
	}
	
	public static void split(String file , int size)
	{		
		
		RandomAccessFile raf;
		long length = 0;
		try {
			raf = new RandomAccessFile(new File(file),"r");
			length = raf.length();
			raf.close();
		} catch (FileNotFoundException e) {
			e.printStackTrace();
		} catch (IOException e) {
			e.printStackTrace();
		}

		for (int i = 0; i < length; i+= size){
			new Thread(new JavaFile(file, i, i + size)).start(); 
		}
		
	}
	
}
