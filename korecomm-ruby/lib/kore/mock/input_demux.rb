module Kore
  module Mock
    class InputDemux
      # Demultiplexes STDIN so for demo purposes, can trigger events
      # from the specified "platform"
      # Ex: "discord !bacon time" or "irc !droplet start awesomesauce"
      def initialize(clients)
        @clients= clients
      end

      def listen
        EM.defer do
          loop do
            input = gets.chomp
            split = input.split(" ")
            client = split.shift

            if !@clients.key?(client.to_sym)
              puts "INPUT DEMUX ERROR: must specify client to send message to, ex: 'discord !bacon time'"
              puts "Available platforms: #{@clients.keys.map{|p| p.to_s}}}"
              next
            end

            identity = "#{client}_duder"
            stripped_input = input[client.length..input.length].strip
            @clients[client.to_sym].call("#{identity}: #{stripped_input}")
          end
        end
      end
    end
  end
end
