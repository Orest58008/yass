package logos

//----------------------------------------//
// CREDITS FOR DISTRO LOGOS AND ASCII ART //
//---------------------------------------//
// Distro logos: ufetch and pfetch       //
// https://gitlab.com/jschx/ufetch       //
// https://github.com/dylanaraps/pfetch  //
//                                       //
// ASCII art: ASCII Art Archive          //
// https://www.asciiart.eu/              //
//---------------------------------------//

var Distros = map[string][]string{
    "alpine": {"<blue>",
	       "<blue>      /\\          ",
	       "<blue>     /  \\         ",
	       "<blue>    / /\\ \\  /\\    ",
	       "<blue>   / /  \\ \\/  \\   ",
	       "<blue>  / /    \\ \\/\\ \\  ",
	       "<blue> / / /|   \\ \\ \\ \\ ",
	       "<blue>/_/ /_|    \\_\\ \\_\\",},
    "antergos": {"<blue>",
		 "<blue>     .```.     ",
		 "<blue>    /     \\    ",
		 "<blue>   /  <white>.`   <blue>\\   ",
		 "<blue>  / <white>.` .` . <blue>\\  ",
		 "<blue> / <white>` .`.` .` <blue>\\ ",
		 "<blue>|    <white>`.`  `   <blue>|",
		 "<blue> \\___________/ ",},
    "arch": {"<cyan>",
	     "<cyan>      /\\      ",
	     "<cyan>     /  \\     ",
	     "<cyan>    /\\   \\    ",
	     "<cyan>   /  __  \\   ",
	     "<cyan>  /  (  )  \\  ",
	     "<cyan> / __|  |__\\\\ ",
	     "<cyan>/.`        `.\\",},
    "archbang": {"<cyan>",
		 "<cyan>      /\\      ",
		 "<cyan>     / _\\_    ",
		 "<cyan>    /  \\ /    ",
		 "<cyan>   /   // \\   ",
		 "<cyan>  /   //   \\  ",
		 "<cyan> / ___()___ \\ ",
		 "<cyan>/.`        `.\\",},
    "arco": {"<blue>",
	     "<blue>      /\\      ",
	     "<blue>     /  \\     ",
	     "<blue>    / /\\ \\    ", 
	     "<blue>   / /  \\ \\   ",
	     "<blue>  / /    \\ \\  ",
	     "<blue> / / _____\\ \\ ",
	     "<blue>/_/  \\_______\\",},
    "artix": {"<cyan>",
	      "<cyan>      /\\      ",
	      "<cyan>     /  \\     ",
	      "<cyan>    /`.. \\    ",
	      "<cyan>   /    `.\\   ",
	      "<cyan>  /   ..`  \\  ",
	      "<cyan> / ..`  `.. \\ ",
	      "<cyan>/.`        `.\\",},
    "centos": {"<green>",
	       "<green> ____<yellow>^<magenta>____ ",
	       "<green> |\\  <yellow>|<magenta>  /| ",
	       "<green> | \\ <yellow>|<magenta> / | ",
	       "<magenta><---- <blue>---->",
	       "<blue> | / <green>|<yellow> \\ | ",
	       "<blue> |/__<green>|<yellow>__\\| ",
	       "<green>     v     ",},
    "crux": {"<cyan>",
	     "<cyan>    ___   ",
	     "<cyan>   (<white>.· <cyan>|  ",
	     "<cyan>   (<yellow><> <cyan>|  ",
	     "<cyan>  / <white>__  <cyan>\\ ",
	     "<cyan> ( <white>/  \\ <cyan>/|",
	     "<yellow>_<cyan>/\\ <white>__)<cyan>/<yellow>_<cyan>)",
	     "<yellow>\\/<cyan>-____<yellow>\\/ ",},
    "debian": {"<red>",
	       "<red>   ,---._ ",
	       "<red> /`  __  \\",
	       "<red>|   /    |",
	       "<red>|   `.__.`",
	       "<red> \\        ",
	       "<red>  `-,_    ",
	       "<red>          ",},
    "devuan": {"<magenta>",
	       "<magenta>-.,          ",
	       "<magenta>   `'-.,     ",
	       "<magenta>        `':. ",
	       "<magenta>           ::",
	       "<magenta>      __--`:`",
	       "<magenta> _,--` _.-`  ",
	       "<magenta>:_,--``      ",},
    "elementary": {"<blue>",
		   "<white>  _______  ",
		   "<white> /  <white>___  <white>\\ ",
		   "<white>/  <white>|  /  /<white>\\",
		   "<white>|<white>__\\_/  /<white> |",
		   "<white>\\   <white>/__/<white>  /  ",
		   "<white> \\_______/  ",
		   "<white>            ",},
    "endeavour": {"<magenta>",
		  "     <red>./<blue>\\     ",
	          "    <red>/<blue>/  \\<cyan>\\.  ",
		  "  <red>./<blue>/    \\ <cyan>\\ ",
		  " <red>/ <blue>/     _) <cyan>)",
	          "<red>/_<blue>/___--`<cyan>__/ ",
	          " <cyan>/____--`    ",},
    "fedora": {"<blue>",
	       "<white>      _____   ",
	       "<white>     /   __)<blue>\\ ",
	       "<white>     |  /  <blue>\\ \\",
	       "<blue>  __<white>_|  |_<blue>_/ /",
	       "<blue> / <white>(_    _)<blue>_/ ",
	       "<blue>/ /  <white>|  |     ",
	       "<blue>\\ \\<white>__/  |     ",
	       "<blue> \\<white>(_____/     ",},
    "gentoo": {"<magenta>",
	       "<magenta>  .-----.    ",
	       "<magenta>.`    _  `.  ",
	       "<magenta>`.   (_)   `.",
	       "<magenta>  `.        /",
	       "<magenta> .`       .` ",
	       "<magenta>/       .`   ",
	       "<magenta>\\____.-`     ",},
    "guix": {"<yellow>",
	     "<yellow>\\____          ____/",
	     "<yellow> \\__ \\        / __/ ",
	     "<yellow>    \\ \\      / /    ",
	     "<yellow>     \\ \\    / /     ",
	     "<yellow>      \\ \\  / /      ",
	     "<yellow>       \\ \\/ /       ",
	     "<yellow>        \\__/        ",},
    "hyperbola": {"<white>",
		  "<white>    /`__.`/   ",
		  "<white>    \\____/    ",
		  "<white>    .--.      ",
		  "<white>   /    \\     ",
		  "<white>  /  ___ \\    ",
		  "<white> / .`   `.\\   ",
		  "<white>/.`      `.\\  ",},
    "instantos": {"<blue>",
		  "<white> ,-''-,     ",
		  "<white>: .''. :    ",
		  "<white>: ',,' :    ",
		  "<white> '-____:__  ",
		  "<white>       :  `.",
		  "<white>       `._.'",},
    "linux": {"<clear>",
	      "<black>    ___   ",
	      "<black>   (<white>.. <black>\\  ",
	      "<black>   (<yellow><> <black>|  ",
	      "<black>  /<white>/  \\ <black>\\ ",
	      "<black> ( <white>|  | <black>/|",
	      "<yellow>_<black>/\\<white>|__)<black>/<yellow>_<black>)",
	      "<yellow>\\/<black>-____<yellow>\\/ ",},
    "linux-lite": {"<yellow>",
		  "<yellow>   /\\ ",
		  "<yellow>  /  \\",
		  "<yellow> / <white>/ <yellow>/",
		  "<yellow>> <white>/ <yellow>/ ",
		  "<yellow>\\ <white>\\ <yellow>\\ ",
		  "<yellow> \\_<white>\\<yellow>_\\",
		  "<white>    \\ ",},
    "linux-mint": {"<green>",
	     "<green> _____________ ",
	     "<green>|_            \\",
	     "<green>  |  <white>| _____  <green>|",
	     "<green>  |  <white>| | | |  <green>|",
	     "<green>  |  <white>| | | |  <green>|",
	     "<green>  |  <white>\\_____/  <green>|",
	     "<green>  \\___________/",},
    "mageia": {"<cyan>",
	       "<cyan>   *    ",
	       "<cyan>    *   ",
	       "<cyan>   **   ",
	       "<white> /\\__/\\ ",
	       "<white>/      \\",
	       "<white>\\      /",
	       "<white> \\____/ ",},
    "manjaro": {"<green>",
		"<green>||||||||| ||||",
		"<green>||||||||| ||||",
		"<green>||||      ||||",
		"<green>|||| |||| ||||",
		"<green>|||| |||| ||||",
		"<green>|||| |||| ||||",
		"<green>|||| |||| ||||",},
    "mx": {"<white>",
	   "<white>    \\\\  /   ",
	   "<white>     \\\\/    ",
	   "<white>      \\\\    ",
	   "<white>   /\\/ \\\\   ",
	   "<white>  /  \\  /\\  ",
	   "<white> /    \\/  \\ ",
	   "<white>/__________\\",},
    "nixos": {"<blue>",
	      "<blue>  \\    \\ //  ",
	      "<blue> ==\\____\\/ //",
	      "<blue>   //    \\// ",
	      "<blue>==//     //==",
	      "<blue> //\\____//   ",
	      "<blue>// /\\   \\== ",
	      "<blue>  // \\   \\  ",},
    "parabola": {"<magenta>",
		 "<magenta>  __ __ __  _  ",
		 "<magenta>.`_//_//_/ / `.",
		 "<magenta>          /  .`",
		 "<magenta>         / .`  ",
		 "<magenta>        /.`    ",
		 "<magenta>       /`      ",
		 "<magenta>               ",},
    "popos": {"<cyan>",
	      "<cyan>______           ",
	      "<cyan>\\   _ \\        __",
	      "<cyan> \\ \\ \\ \\      / /",
	      "<cyan>  \\ \\_\\ \\    / / ",
	      "<cyan>   \\  ___\\  /_/  ",
	      "<cyan>    \\ \\    _     ",
	      "<cyan>   __\\_\\__(_)_   ",
	      "<cyan>  (___________)  ",},
    "pureos": {"<clear>",
	       "<clear> _____________ ",
	       "<clear>|  _________  |",
	       "<clear>| |         | |",
	       "<clear>| |         | |",
	       "<clear>| |_________| |",
	       "<clear>|_____________|",},
    "raspbian": {"<red>",
		      "<green>  __  __  ",
		      "<green> (_\\)(/_) ",
		      "<red> (_(__)_) ",
		      "<red>(_(_)(_)_)",
		      "<red> (_(__)_) ",
		      "<red>   (__)   ",},
    "slackware": {"<blue>",
		  "<blue>   ________  ",
		  "<blue>  /  ______| ",
		  "<blue>  | |______  ",
		  "<blue>  \\______  \\ ",
		  "<blue>   ______| | ",
		  "<blue>| |________/ ",
		  "<blue>|____________",},
    "solus": {"<blue>",
	      "<white>    /|     ",
	      "<white>   / |\\    ",
	      "<white>  /  | \\__ ",
	      "<white> /<blue>___<white>|<blue>__<white>\\<blue>_<white>\\",
	      "<blue>\\         /",
	      "<blue> `-------' ",},
    "suse": {"<green>",
	     "<green>    _______  ",
	     "<green>___|    __ \\",
	     "<green>       / <white>.<green>\\ \\",
	     "<green>       \\__/ <green>|",
	     "<green>     _______|",
	     "<green>     \\______ ",
	     "<green>-__________/",},
    "ubuntu": {"<red>",
	       "<red>         _  ",
	       "<red>     ---(_) ",
	       "<red> _/  ---  \\ ",
	       "<red>(_) |   |   ",
	       "<red>  \\  --- _/ ",
	       "<red>     ---(_) ",
	       "<red>            ",},
    "void": {"<green>",
	     "<green>    _______    ",
	     "<green>    \\_____ `-  ",
	     "<green> /\\   ___ `- \\ ",
	     "<green>| |  /   \\  | |",
	     "<green>| |  \\___/  | |",
	     "<green> \\ `-_____  \\/ ",
	     "<green>  `-______\\    ",},
    "voyager": {"<yellow>",
		"<yellow>   _____   ____ ",
		"<yellow>  |     | |    |",
		"<yellow>  |     | |    |",
		"<yellow>  |     | |    |",
		"<yellow>  |     | |____|",
		"<yellow>  |     |______ ",
		"<yellow>  |            |",
		"<yellow>  |____________|",},
    //easter eggs
    "invader": {"<magenta>",
		"<magenta>      ____               ",
		"<magenta>     /___/\\__            ",
		"<magenta>    _\\   \\/_/\\__         ",
		"<magenta>  __\\       \\/_/\\        ",
		"<magenta>  \\   __    __ \\ \\       ",
		"<magenta> __\\  \\_\\   \\_\\ \\ \\   __ ",
		"<magenta>/_/\\\\   __   __  \\ \\_/_/\\",
		"<magenta>\\_\\/_\\__\\/\\__\\/\\__\\/_\\_\\/",
		"<magenta>   \\_\\/_/\\       /_\\_\\/  ",
		"<magenta>      \\_\\/       \\_\\/    ",},
    "kitty": {"<clear>",
	      "<clear>      /\\_/\\ ",
	      "<clear> /\\  / <blue>o o<clear> \\",
	      "<clear>//\\\\ \\~(*)~/",
	      "<clear>`  \\/   ^ / ",
	      "<clear>   | \\|| || ",
	      "<clear>   \\ '|| || ",
	      "<clear>    \\)()-())",},
    "maple-leaf": {"<red>",
		   "<red>     .\\^/.     ",
		   "<red>   . |`|/| .   ",
		   "<red>   |\\|\\|'|/|   ",
		   "<red>.--'-\\`|/-''--.",
		   "<red> \\`-._\\|./.-'/ ",
		   "<red>  >`-._|/.-'<  ",
		   "<red> '~|/~~|~~\\|~' ",
		   "<red>       |       ",},
    "totoro": {"<white>",
	       "<magenta> _____                           ",
	       "<magenta>/     \\                          ",
	       "<magenta>vvvvvvv<black>  /|__/|                  ",
	       "   I   <black>/ <white>O<black>,<white>O<black>  |                  ",
	       "   I <black>/<white>_____   <black>|      <blue>/|/|        ",
	       "   J<white>/^ ^ ^ \\<white>  <black>|    <blue>/<white>00<white>  <blue>|    <white>_//|",
	       "   <white>|^ ^ ^ ^ |<black>W|   <blue>|<white>/<blue>^^<white>\\ <blue>|   <white>/oo |",
	       "    <white>\\<black>m<white>___<black>m<white>__|<black>_|    <white>\\<blue>m<white>_<blue>m_|   <white>\\mm_|",},
    "unicorn": {"<magenta>",
		"<white>        <magenta>_.%%%%%%%%<white>(/_______.",
		"<white>   _.-<magenta>'% % % % %<white>   ;-'-'''` ",
		"<white>.-'            / (<magenta>q<white>:        ",
		"<white> '.    '.    '(__: :        ",
		"<white>     '.   '.'  \\::.:        ",
		"<white>       ;    :   \\i_o        ",
		"<white>       '.   :               ",
		"<white>         `''                ",},
}
