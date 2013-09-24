package com.example.util;


import java.io.*;
import java.net.*;


 public class SocketClient {
    private Socket client;
    private String site;
    private Integer port;
    
    public SocketClient(String site, int port){
    	this.site = site;
    	this.port = port;
    }
    
    /**
     * 注册socket
     * @date 2013年7月5日 下午5:06:18
     * @author LiLin
     * @email Lvan_software@foxmail.com
     * 
     * @return
     * @return boolean
     */
    public boolean initSocket(){
    	boolean tag = false;
    	try{
//    		client = new Socket(site,port);
    		
    		client = new Socket();  
    		InetSocketAddress isa = new InetSocketAddress(site,port);  
    		client.connect(isa,100);
    		
    		tag = true;
//    		System.out.println("Client is created! site:"+site+" port:"+port);
    	}catch (UnknownHostException e){
    		e.printStackTrace();
    	}catch (IOException e){
    		e.printStackTrace();
    	}catch (Exception e){
    		e.printStackTrace();
    	}
    	return tag;
    }
    
    /**
     * 检测socket是否有效
     * @date 2013年7月5日 下午4:57:06
     * @author LiLin
     * @email Lvan_software@foxmail.com
     * 
     * @return
     * @return boolean
     */
    public boolean isConnected(){
    	if (client != null && client.isConnected()){
    		return true;
    	}else{
    		return false;
    	}
    }
    
    /**
     * 发送消息
     * @date 2013年7月5日 下午5:06:00
     * @author LiLin
     * @email Lvan_software@foxmail.com
     * 
     * @param msg
     * @return
     * @return boolean
     */
    public boolean sendMsg(String msg){
        try{
        	
//            BufferedReader in = new BufferedReader(new InputStreamReader(client.getInputStream()));
//            PrintWriter out = new PrintWriter(client.getOutputStream());
//            out.println(msg);
//            out.flush();
//            return in.readLine();
            
            BufferedWriter writer = new BufferedWriter(new OutputStreamWriter(client.getOutputStream()));
            writer.write(msg);
            writer.flush();
            writer.close();
            writer = null;
            
            return true;
        }catch(IOException e){
            e.printStackTrace();
            return false;
        }
    }
    public void closeSocket(){
        try{
        	if (client != null && client.isConnected()){
        		client.close();
        	}
        }catch(IOException e){
            e.printStackTrace();
        }
    }
    public static void main(String[] args) throws Exception{
    	SocketClient socket = new SocketClient("127.0.0.1", 9090);
    	socket.initSocket();
    	socket.sendMsg("123456789");
    	socket.closeSocket();
    }

}
