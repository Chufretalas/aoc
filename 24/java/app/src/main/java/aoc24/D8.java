package aoc24;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.Optional;
import java.util.Set;

record Node(int x, int y, char nodeType) {
}

public class D8 {
    static public void part1() {
        ArrayList<Node> antennas = new ArrayList<>();
        Set<Node> antinodes = new HashSet<Node>();
        int maxX = 0;
        int maxY = 0;

        try (LineReader reader = new LineReader("d8.txt")) {
            Optional<String> lineOption = reader.nextLine();

            while (lineOption.isPresent()) {
                String line = lineOption.get();
                maxX = line.length();
                maxY++;
                char[] chars = line.toCharArray();
                for (int i = 0; i < chars.length; i++) {
                    if (chars[i] != '.') {
                        antennas.add(new Node(i, maxY - 1, chars[i]));
                    }
                }
                lineOption = reader.nextLine();
            }
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }

        for (Node currentNode : antennas) {
            for (Node otherNode : antennas) {
                if (currentNode.equals(otherNode) || currentNode.nodeType() != otherNode.nodeType())
                    continue;

                // looking at the mirroed position of the otherNode using currentNode as the
                // mirror
                // current - other = a vector pointing to current from other
                // add this vector to current to position it correctly
                int antinodeX = 2 * currentNode.x() - otherNode.x();
                int antinodeY = 2 * currentNode.y() - otherNode.y();

                // checking if the antinode would be out of bounds
                if (antinodeX < 0 || antinodeX >= maxX || antinodeY < 0 || antinodeY >= maxY)
                    continue;

                antinodes.add(new Node(antinodeX, antinodeY, '#'));
            }
        }

        System.err.println(antinodes.size());
    }
}
