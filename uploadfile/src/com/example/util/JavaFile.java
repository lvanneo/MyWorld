package com.example.util;

import java.io.File;
import java.io.IOException;
import java.io.RandomAccessFile;

public class JavaFile implements Runnable{
	
	private String data;
	private String fileName;
	private long begin;
	private long end;
	private int size;
	private int seqnum;
	private String mergeFileName;
	
	
	public JavaFile( int size, int seqnum, String file, String mergeFileName, long begin,long end){
		this.fileName = file;
		this.begin = begin;
		this.end = end;
		this.size = size;
		this.seqnum = seqnum;
		this.mergeFileName = mergeFileName;
	}
	
	@Override
	public void run() {
		
		write(fileName, begin, end);
		
		SocketClient socket = new SocketClient("192.168.1.5", 9090);
    	socket.initSocket();
    	socket.sendMsg(data);
    	socket.closeSocket();
	}
	
	
	private long write(String file,long begin,long end) {
		SocketClient socket = new SocketClient("192.168.1.5", 9090);
    	socket.initSocket();
    	
    	char cc = (char)seqnum;
    	String seq = "" + cc + mergeFileName;
    	socket.sendMsg(seq);
    	
		RandomAccessFile in;
		long endPointer = 0;
		try {
			in = new RandomAccessFile(new File(file),"r");
			byte[] bb = new byte[1024 * size];
			in.seek(begin);//从指定位置读取
			
			StringBuffer data = new StringBuffer();
			while(in.getFilePointer() < in.length() && in.getFilePointer() <= end && in.read(bb)!= -1)
			{
//				data.append(new String(bb));
				socket.sendMsg(new String(bb));
			}
			this.data += data.toString();
			
			endPointer =in.getFilePointer();
			in.close();
		} catch (IOException e) {
			e.printStackTrace();
		}
		socket.closeSocket();
		return endPointer;
	}
	
}
