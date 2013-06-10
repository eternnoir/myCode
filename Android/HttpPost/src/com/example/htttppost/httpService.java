package com.example.htttppost;
import java.io.IOException;
import java.util.List;

import org.apache.http.HttpEntity;
import org.apache.http.HttpResponse;
import org.apache.http.NameValuePair;
import org.apache.http.ParseException;
import org.apache.http.client.HttpClient;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.DefaultHttpClient;
import org.apache.http.util.EntityUtils;

public class httpService {

	public String httpGet(String url,String value){
		String _ret = new String();
		HttpClient client = new DefaultHttpClient();
		HttpGet get = new HttpGet(url+"?"+value);
		try {
			HttpResponse response = client.execute(get);
			HttpEntity resEntity = response.getEntity();
			_ret = EntityUtils.toString(resEntity);
		} catch (ParseException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			return "ERROR";
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			return "ERROR";
		}
		catch (Exception e){
			e.printStackTrace();
			return"NOooooo \n check your url";
		}
		return _ret;
	}
}
