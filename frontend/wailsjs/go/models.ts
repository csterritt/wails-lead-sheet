export namespace parser {
	
	export class ParsedContent {
	
	
	    static createFrom(source: any = {}) {
	        return new ParsedContent(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

