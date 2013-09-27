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
	
	public static void main(String[] args) {
		byte[] bb = {'r','h','m'};
		String ss = new String(bb);
		System.out.println(ss);
		
		char cc = (char)6;
		String str = "" + cc;
		System.out.println(str);
	}
	
	public static void main2(String[] args) throws Exception {
		Test.split("F:\\Temp\\wpfImage.zip", "F:\\Temp\\wpfImage22.zip", 10, 1024 * 1024);
	}
	
	public static void split(String file , String mergeFileName, int xiannum , int size)
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
		
		int datasize = (int)length / xiannum;
		
		int seatNum = 0;
		for (int i = 0; i < xiannum; i++){
			if (i == xiannum){
				new Thread(new JavaFile( size, i, file, mergeFileName, seatNum, length)).start(); 
			}else{
				new Thread(new JavaFile( size, i, file, mergeFileName, seatNum, i + datasize)).start(); 
			}
			seatNum += datasize;
		}
		
		SocketClient socket = new SocketClient("192.168.1.5", 9090);
    	socket.initSocket();
    	char cxian = (char)xiannum;
    	socket.sendMsg("fileover" + cxian + mergeFileName);
    	socket.closeSocket();
		
	}
	
}
