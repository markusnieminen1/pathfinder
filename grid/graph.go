package grid

import (
	"fmt"
	"os"
	"pathfinder/data"
)

/*
AI GENERATED CODE.

HAVEN'T VERIFIED ANYTHING OR EVEN LOOKED AT IT.

JUST TO GET AN IDEA OF THE GRAPHS.

*/

func ExportTopologyHTML(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprint(f, `<!DOCTYPE html>
<html>
<head>
    <title>Network Topology Coordinate Map</title>
    <style>
        body { margin: 0; background-color: #111216; font-family: sans-serif; overflow: hidden; color: #fff; }
        #title { position: absolute; top: 20px; left: 20px; pointer-events: none; }
        h1 { margin: 0; font-size: 18px; }
        p { margin: 5px 0 0 0; color: #8a8f98; font-size: 12px; }
        canvas { display: block; width: 100vw; height: 100vh; }
    </style>
</head>
<body>
    <div id="title">
        <h1>Transit Grid Map</h1>
        <p>Locked by Coordinates (Scroll to zoom, Drag background to pan)</p>
    </div>
    <canvas id="graphCanvas"></canvas>

    <script>
        const canvas = document.getElementById('graphCanvas');
        const ctx = canvas.getContext('2d');
        
        let width = canvas.width = window.innerWidth;
        let height = canvas.height = window.innerHeight;

        // Camera Pan & Zoom state variables
        let offsetX = width / 2;
        let offsetY = height / 2;
        let scale = 7.0; // Multiplier to blow up small (0-100) coordinates into screen pixels
        let isDragging = false;
        let startX, startY;

        window.addEventListener('resize', () => {
            width = canvas.width = window.innerWidth;
            height = canvas.height = window.innerHeight;
            render();
        });

        // 1. Inject Stations with fixed coordinates directly from Go
        const nodes = [
`)

	// Loop over map and write out fixed coordinates
	for name, station := range data.StationsMap {
		// Flip or offset the Y calculation if your coordinates assume 0,0 is bottom-left
		fmt.Fprintf(f, "            { id: %q, cx: %d, cy: %d, r: 5 },\n", name, station.Coordinates[0], station.Coordinates[1])
	}

	fmt.Fprint(f, "        ];\n\n        // 2. Inject connections\n        const links = [\n")

	// Inject structural linkages uniquely
	seen := make(map[string]bool)
	for name, station := range data.StationsMap {
		for _, conn := range station.Connections {
			key1 := name + "|||" + conn.Name
			key2 := conn.Name + "|||" + name
			if !seen[key1] && !seen[key2] {
				fmt.Fprintf(f, "            { source: %q, target: %q },\n", name, conn.Name)
				seen[key1] = true
			}
		}
	}

	fmt.Fprint(f, `        ];

        // Map node IDs to objects for fast rendering lookups
        const nodeMap = {};
        nodes.forEach(n => nodeMap[n.id] = n);
        links.forEach(l => {
            l.sourceObj = nodeMap[l.source];
            l.targetObj = nodeMap[l.target];
        });

        // Auto-center camera on the average coordinate node cluster at launch
        if (nodes.length > 0) {
            let sumX = 0, sumY = 0;
            nodes.forEach(n => { sumX += n.cx; sumY += n.cy; });
            const avgX = sumX / nodes.length;
            const avgY = sumY / nodes.length;
            
            // Set offsets so the cluster midpoint maps to screen center
            offsetX = width / 2 - (avgX * scale);
            offsetY = height / 2 - (-avgY * scale); // Negative because canvas Y space runs downwards
        }

        // --- Interaction Handlers (Pan and Zoom) ---
        canvas.addEventListener('mousedown', e => {
            isDragging = true;
            startX = e.clientX - offsetX;
            startY = e.clientY - offsetY;
        });

        canvas.addEventListener('mousemove', e => {
            if (isDragging) {
                offsetX = e.clientX - startX;
                offsetY = e.clientY - startY;
                render();
            }
        });

        window.addEventListener('mouseup', () => isDragging = false);

        canvas.addEventListener('wheel', e => {
            e.preventDefault();
            const zoomFactor = 1.1;
            
            // Track pointer position to zoom relative to your mouse cursor position
            const mouseX = e.clientX;
            const mouseY = e.clientY;
            const graphX = (mouseX - offsetX) / scale;
            const graphY = (mouseY - offsetY) / scale;

            if (e.deltaY < 0) {
                scale *= zoomFactor;
            } else {
                scale /= zoomFactor;
            }

            offsetX = mouseX - graphX * scale;
            offsetY = mouseY - graphY * scale;
            render();
        }, { passive: false });

        // Transform formulas mapping raw coordinate bounds to camera space viewport
        function getScreenX(graphX) { return graphX * scale + offsetX; }
        function getScreenY(graphY) { return -graphY * scale + offsetY; } // Inverts Y axis so higher Y coordinate goes UP

        function render() {
            ctx.clearRect(0, 0, width, height);

            // Draw Transit Line Connections
            ctx.strokeStyle = '#334155';
            ctx.lineWidth = 1.5;
            links.forEach(l => {
                if (!l.sourceObj || !l.targetObj) return;
                ctx.beginPath();
                ctx.moveTo(getScreenX(l.sourceObj.cx), getScreenY(l.sourceObj.cy));
                ctx.lineTo(getScreenX(l.targetObj.cx), getScreenY(l.targetObj.cy));
                ctx.stroke();
            });

            // Draw Station Points & Labels
            nodes.forEach(n => {
                const sx = getScreenX(n.cx);
                const sy = getScreenY(n.cy);

                // Draw solid layout node points
                ctx.beginPath();
                ctx.arc(sx, sy, n.r, 0, 2 * Math.PI);
                ctx.fillStyle = '#3b82f6';
                ctx.fill();
                ctx.strokeStyle = '#ffffff';
                ctx.lineWidth = 1;
                ctx.stroke();

                // Draw Text identifiers next to the node point
                ctx.fillStyle = '#cbd5e1';
                ctx.font = '10px monospace';
                ctx.fillText(n.id, sx + 9, sy + 3);
            });
        }
        
        render();
    </script>
</body>
</html>`)

	return nil
}
