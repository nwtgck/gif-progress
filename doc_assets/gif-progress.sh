# !/bin/bash
function continueRequest(){
  if [[ $1 =~ [^\\]+\.gif+$ ]]; then
   read -p 'Please enter the final path for the gif: ' finalPath
   read -p 'Enter the height of the progress bar [5]: ' height
   size=${height:-5}
   if [[ "$size" -le "0" ]]; then
     echo "The height must be higher than 0"
   else 
    read -p "Write the hex code for the color of the progress bar [#ccc]: " hex
    finalHex=${hex:-#ccc}
    if [[ $finalHex =~ ^#([0-9a-fA-F]{3}){1,2}$ ]]; then
       echo "Answer 'y' or 'n'."
       read -p "You want the progress bar at the top of the file? [n]: " top
       echo "Generating new gif file..."
       shouldBeTop=${top:-n}
       if [[ "$shouldBeTop" == "y" ]]; then
         cat $1 | gif-progress --bar-color=${hex:-#ccc} --bar-height=${height:-5} --bar-top > "$finalPath"
       else
         cat $1 | gif-progress --bar-color=${hex:-#ccc} --bar-height=${height:-5} > "$finalPath"
       fi
    else
       echo "The hex code '$hex' is invalid. Please enter a valid hex code!"
    fi
   fi
  else
   echo "Invalid gif file! Please specify a valid path to a gif file"
  fi
}


read -p "Type in the path to the gif: " path

if [ -d "$path" ]; then
  echo "The path $path is a directory. You must specify a file!"
  exit -1
elif [ -f "$path" ]; then
  continueRequest "$path"
else
  echo "Invalid path! Please specify a path to the gif"
  exit -1
fi
