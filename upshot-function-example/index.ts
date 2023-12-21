import { Console } from "as-wasi/assembly";
import { CgiCommand, cgiExtendsList} from "@blockless/sdk/assembly/cgi";
import { buffer2string} from "@blockless/sdk/assembly/strings";

// get as list of the extensions available
let l = cgiExtendsList();

if (l != null) {
    let command = new CgiCommand("cgi-upshot", null, null);
    let rs = command.exec();
    if (rs == true) {
        const SEP = "\r\n";
        let buf = new Array<u8>(1024);
        let req = '{"field1":"foo"}';
        let req_len = req.length;
        let head = `${req_len}${SEP}`;
        command.stdinWriteString(head);
        command.stdinWriteString(`${req}${SEP}`);
        buf = new Array<u8>(65535);
        let all_buff: u8[] = new Array(0);
        let l = command.stdoutRead(buf);
        all_buff = all_buff.concat(buf.slice(0, l));
        let read_string = buffer2string(all_buff, all_buff.length);
        Console.log(read_string);
    }
    command.close();
}