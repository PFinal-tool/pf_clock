export namespace main {
	
	export class Config {
	    clock_bg_color: string;
	    clock_text_color: string;
	    text_bg_color: string;
	    text_color: string;
	    text_content: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.clock_bg_color = source["clock_bg_color"];
	        this.clock_text_color = source["clock_text_color"];
	        this.text_bg_color = source["text_bg_color"];
	        this.text_color = source["text_color"];
	        this.text_content = source["text_content"];
	    }
	}

}

