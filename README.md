
# bredit


bredit stands for: **B**encode **R**eplacer and **EDIT**or  - which is written in the Golang language.

It's released under GPLv2 license and therefor is free software.

bredit is used to edit and replace data that utilizes the bencode format such as torrent & rtorrent files.
It's designed to edit files in bulk but you can manually edit a file if needed.

Currently bredit only supports CLI but is planned to support editing via GUI.


# Usage
in CLI mode you need to give bredit at least 2 flags and 1 arg:
    

 - -k --- takes a string that represents a __key__ in the bencode dictionary
 - -val --- takes a string that represents the value used to replace the **key**'s value
 - args[0] --- takes a string that represents a full path to either a file or folder

#### Example:

- Single file:
 

       bredit -k="encoding" -val="UTF-8" /home/user/.session/Alice\ in\ Wonderland.torrent
- Folder:
 

       bredit -k="encoding" -val="SHIFT-JIS" /home/user/.session/
      
---
There are some limitations currently when using bredit in CLI mode.
