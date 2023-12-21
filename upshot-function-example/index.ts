import { Console } from "as-wasi/assembly";
import { CGIExtension, CgiCommand, Env, cgiExtendsList} from "@blockless/sdk/assembly/cgi";
import { buffer2string, string2buffer, arrayIndex} from "@blockless/sdk/assembly/strings";
import { Arr } from "@blockless/sdk/assembly/json/JSON";


// get as list of the extensions available
let l = cgiExtendsList();

// extension did not register on run
if (l != null) {
    Console.log(`${(l as CGIExtension[]).toString()}`);
}

// let command = new CgiCommand("cgi-upshot", null, null);
// let rs = command.exec();

