package com.example.util;

import java.io.File;
import java.io.IOException;
import java.io.RandomAccessFile;

public class JavaFile implements Runnable{
	
	private String data;
	private String fileName;
	private long begin;
	private long end;
	
	
	public JavaFile(String file,long begin,long end){
		this.fileName = file;
		this.begin = begin;
		this.end = end;
	}
	
	@Override
	public void run() {
		write(fileName, begin, end);
		SocketClient socket = new SocketClient("192.168.1.9", 9090);
    	socket.initSocket();
    	socket.sendMsg(data);
    	socket.closeSocket();
	}
	
	
	private long write(String file,long begin,long end) {
		RandomAccessFile in;
		long endPointer = 0;
		try {
			in = new RandomAccessFile(new File(file),"r");
			byte[] b = new byte[1024];
			in.seek(begin);//从指定位置读取
			
			StringBuffer data = new StringBuffer();
			while(in.getFilePointer() < in.length() && in.getFilePointer() <= end && in.read(b)!= -1)
			{
				data.append(new String(b));
			}
			this.data += data.toString();
			
			endPointer =in.getFilePointer();
			in.close();
		} catch (IOException e) {
			e.printStackTrace();
		}
		return endPointer;
	}
	
}
