package aoc24;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.Optional;

public class LineReader implements AutoCloseable {

    private final BufferedReader reader;

    public LineReader(String fileName) {
        InputStream stream = ClassLoader.getSystemClassLoader().getResourceAsStream(fileName);
        this.reader = new BufferedReader(new InputStreamReader(stream));
    }

    public Optional<String> nextLine() throws IOException {
        String line = this.reader.readLine();
        return Optional.ofNullable(line);
    }

    @Override
    public void close() throws IOException {
        this.reader.close();
    }

}
