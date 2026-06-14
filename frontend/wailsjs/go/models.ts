export namespace main {
	
	export class RedactionEvent {
	    ruleName: string;
	    originalText: string;
	    replacement: string;
	    startIndex: number;
	    endIndex: number;
	
	    static createFrom(source: any = {}) {
	        return new RedactionEvent(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ruleName = source["ruleName"];
	        this.originalText = source["originalText"];
	        this.replacement = source["replacement"];
	        this.startIndex = source["startIndex"];
	        this.endIndex = source["endIndex"];
	    }
	}
	export class RedactResult {
	    redactedText: string;
	    events: RedactionEvent[];
	
	    static createFrom(source: any = {}) {
	        return new RedactResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.redactedText = source["redactedText"];
	        this.events = this.convertValues(source["events"], RedactionEvent);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class RedactionRule {
	    id: string;
	    name: string;
	    pattern: string;
	    isEnabled: boolean;
	    isLiteral: boolean;
	    category: string;
	    source: string;
	    severity: string;
	    description: string;
	    replacement: string;
	
	    static createFrom(source: any = {}) {
	        return new RedactionRule(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.pattern = source["pattern"];
	        this.isEnabled = source["isEnabled"];
	        this.isLiteral = source["isLiteral"];
	        this.category = source["category"];
	        this.source = source["source"];
	        this.severity = source["severity"];
	        this.description = source["description"];
	        this.replacement = source["replacement"];
	    }
	}

}

