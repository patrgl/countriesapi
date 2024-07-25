<h1>Description</h1>

I made a simple API that returns country information in an organized JSON format in order to practice Golang. <br>
This is meant to be used as an example and a starting off point.

<h1>Examples: </h1>

<b>Path:</b> /countries/name/pan <br>
<b>Returns:</b> <br>
&emsp;{<br>
&emsp;"status": "success",<br>
&emsp;"statusCode": 200,<br>
&emsp;"data": [<br>
&emsp;{<br>
      &emsp;&emsp;"name": "Japan",<br>
      &emsp;&emsp;"capital": "Tokyo",<br>
      &emsp;&emsp;"phone_code": "81",<br>
      &emsp;&emsp;"currency_name": "Japanese yen",<br>
      &emsp;&emsp;"region": "Asia"<br>
&emsp;},<br>
&emsp;{<br>
      &emsp;&emsp;"name": "Panama",<br>
      &emsp;&emsp;"capital": "Panama City",<br>
      &emsp;&emsp;"phone_code": "507",<br>
      &emsp;&emsp;"currency_name": "Panamanian balboa",<br>
      &emsp;&emsp;"region": "Americas"<br>
&emsp;}<br>
&emsp;]<br>
&emsp;}<br>
<br>
<b>Path:</b> /countries/phonecode/56<br>
<b>Returns:</b> <br>
&emsp;{<br>
  &emsp;"status": "success",<br>
  &emsp;"statusCode": 200,<br>
  &emsp;"data": [<br>
    &emsp;{<br>
      &emsp;&emsp;"name": "Chile",<br>
      &emsp;&emsp;"capital": "Santiago",<br>
      &emsp;&emsp;"phone_code": "56",<br>
      &emsp;&emsp;"currency_name": "Chilean peso",<br>
      &emsp;&emsp;"region": "Americas"<br>
    &emsp;}<br>
  &emsp;]<br>
&emsp;}<br>
<br>
<br>
<br>
I got the DB from: https://github.com/dr5hn/countries-states-cities-database?tab=readme-ov-file <br>
